GO111MODULES=on
APP:=$(notdir $(shell pwd))
# DOCKER_USERNAME:=$(shell docker info | sed '/Username:/!d;s/.* //')
# Use organization name instead
DOCKER_USERNAME:=fluoride
DOCKER_REGISTRY?=$(shell docker info | sed '/Registry:/!d;s/.* //')
COMMIT_SHA=$(shell git rev-parse --short HEAD)
DOCKER_ENV?=$(shell env | grep "DOCKER")

default: build

.PHONY: docs
## docs: generates OpenAPI 2.0 docs and API Blueprint equivalents
docs:
	@echo "Generating OpenAPI 2.0 spec"
	@swagger generate spec -m -o docs/openapi/oas-2.yaml
	@echo "Enriching OpenAPI 2.0 spec with code snippets"
	@openapi-snippet docs/openapi/oas-2.yaml -o docs/openapi/oas-2.yaml -t c -t csharp -t go -t java -t javascript -t node -t php -t python -t ruby -t shell -t swift
	@echo "Generating OpenAPI 3.0 spec from OpenAPI 2.0 spec"
	@api-spec-converter --from=swagger_2 --to=openapi_3 --syntax=yaml docs/openapi/oas-2.yaml > docs/openapi/oas-3.yaml
	@echo "Generating API Blueprint from generated OpenAPI 2.0 spec"
	@fury --format text/vnd.apiblueprint docs/openapi/oas-2.yaml docs/apib/api.apib
	


.PHONY: docs-serve
## docs-serve: serves swagger docs with redoc flavor
docs-serve: docs
	swagger serve -F=redoc docs/openapi/oas-2.yaml

.PHONY: dependencies
## dependencies: download dependencies
dependencies: clean
	@echo "Downloading dependencies"
	@go mod download

.PHONY: install
## install: installs dependencies
install:
	go install -v

.PHONY: build
## build: build the application
build: clean
	@echo "Building binary"
	@CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -tags netgo -ldflags "-s -w" -o bin/${APP}
	@echo "Binary built (size `stat -c '%s' "bin/${APP}" | numfmt --to=si --suffix=B`)"

.PHONY: run
## run: runs go run main.go
run: build
	@echo
	@./bin/${APP}

.PHONY: serve
## serve: runs go run main.go serve
serve: build
	@echo
	@./bin/${APP} serve

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning"
	@echo "Deleting binaries"
	@rm -rf bin
	@go clean
	@echo "Tidying go mod"
	@go mod tidy

.PHONY: compose
## compose: docker compose build & up
compose: compose-build compose-up

.PHONY: compose-clean
## compose-clean: shuts down and removes the containers
compose-clean:
	@sudo docker-compose down && \
	sudo docker-compose rm

.PHONY: compose-build
## compose-build: builds using docker-compose
compose-build: compose-clean
	@sudo ${DOCKER_ENV} docker-compose build --build-arg TARGETOS=${TARGETOS} --build-arg TARGETARCH=${TARGETARCH}

.PHONY: compose-up
## compose-up: runs using docker compose
compose-up: compose-build
	@sudo docker-compose up

.PHONY: docker-build-server
## docker-build-server: builds the docker image for the server
docker-build-server:
	@sudo ${DOCKER_ENV} docker build -t ${DOCKER_USERNAME}/server:${COMMIT_SHA} -f docker/server/Dockerfile .

.PHONY: docker-build-postgres
## docker-build-postgres: build db Dockerfile
docker-build-postgres:
	@sudo ${DOCKER_ENV} docker build -t ${DOCKER_USERNAME}/postgres:${COMMIT_SHA} -f docker/postgres/Dockerfile .

.PHONY: docker-build
## docker-build: builds both server and db Dockerfiles
docker-build: docker-build-server docker-build-postgres

.PHONY: docker-login
## docker-login: pushes the server and postgres docker images to DOCKER_REGISTRY
docker-login:
	@sudo docker login


.PHONY: docker-push
## docker-push: pushes the server and postgres docker images to DOCKER_REGISTRY
docker-push: docker-login docker-push-postgres docker-push-server

.PHONY: docker-push-postgres
## docker-push-postgres: pushes the db docker image to DOCKER_REGISTRY
docker-push-postgres: docker-build-postgres
	@sudo docker push ${DOCKER_USERNAME}/postgres:${COMMIT_SHA}

.PHONY: docker-push-server
## docker-push-server: pushes the server docker image to DOCKER_REGISTRY
docker-push-server: docker-build
	
	@sudo docker push ${DOCKER_USERNAME}/server:${COMMIT_SHA}

.PHONY: help
## help: Prints this help message
help:
	@echo -e "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
