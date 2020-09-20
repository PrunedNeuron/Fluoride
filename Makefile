GO111MODULES=on
APP?=$(notdir $(shell pwd))
REGISTRY?=$(shell docker info | sed '/Registry:/!d;s/.* //')
DOCKER_USERNAME=$(shell docker info | sed '/Username:/!d;s/.* //')
COMMIT_SHA=$(shell git rev-parse --short HEAD)
TARGETOS=linux
TARGETARCH=amd64
ENV=staging

default: build

.PHONY: build
## build: build the application
build: clean
	@echo "Building binary"
	@CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -tags netgo -ldflags "-s -w" -o bin/${APP}
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

# helper rule for deployment
check-environment:
ifndef ENV
	$(error ENV not set, allowed values - `staging` or `production`)
endif

.PHONY: compose
## compose: docker compose build & up
compose: compose-build compose-up

.PHONY: compose-clean
## compose-clean: shuts down and removes the containers
compose-clean:
	sudo docker-compose down
	sudo docker-compose rm

.PHONY: compose-build
## compose-build: builds using docker-compose
compose-build: compose-clean
	sudo docker-compose build --build-arg TARGETOS=${TARGETOS} --build-arg TARGETARCH=${TARGETARCH}

.PHONY: compose-up
## docker-compose: builds using docker-compose
compose-up: compose-build
	sudo docker-compose up

.PHONY: docker-build-server
## docker-build-server: builds the docker image for the server
docker-build-server:
	sudo docker build -t ${DOCKER_USERNAME}/${APP}:${COMMIT_SHA} -f docker/server/Dockerfile .

.PHONY: docker-build-db
## docker-build-db: build db Dockerfile
docker-build-db: 
	sudo docker build -t ${DOCKER_USERNAME}/${APP}:${COMMIT_SHA} -f docker/db/Dockerfile .

.PHONY: docker-build
## docker-build: builds both server and db Dockerfiles
docker-build: docker-build-server docker-build-db

.PHONY: docker-push
## docker-push: pushes the docker image to registry
docker-push: check-environment docker-build
# sudo docker push ${REGISTRY}/${DOCKER_USERNAME}/${APP}:${COMMIT_SHA}
	sudo docker login
	sudo docker push ${DOCKER_USERNAME}/${APP}:${COMMIT_SHA}

.PHONY: help
## help: Prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
