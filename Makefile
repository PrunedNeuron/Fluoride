# EXECUTABLE=Fluoride

# .PHONY: all build build-static

# default: build build-static

# build: ## Builds a dynamic linked binary
# 	go version
# 	go build -o bin/${EXECUTABLE}

# build-static: ## Builds a static binary
# 	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o bin/static/${EXECUTABLE}

GO111MODULES=on
APP?=fluoride
REGISTRY?=gcr.io/images
COMMIT_SHA=$(shell git rev-parse --short HEAD)
TARGETOS=linux
TARGETARCH=amd64

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
	# go run main.go
	./bin/fluoride

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning"
	@go clean
	@echo "Tidying go mod"
	@go mod tidy

.PHONY: test
## test: runs go test with default values
test:
	go test -v -count=1 -race ./...

# helper rule for deployment
check-environment:
ifndef ENV
	$(error ENV not set, allowed values - `staging` or `production`)
endif

.PHONY: docker-build-server
## docker-build-server: builds the docker image for the server
docker-build-server:
	sudo docker build -t ${APP}:${COMMIT_SHA} -f docker/server/Dockerfile .

.PHONY: docker-compose-build
## docker-compose-build: builds using docker-compose
docker-compose-build:
	sudo docker-compose build --build-arg TARGETOS=linux --build-arg TARGETARCH=amd64

.PHONY: docker-compose-up
## docker-compose: builds using docker-compose
docker-compose-up: docker-compose-build
	sudo docker-compose up

.PHONY: docker-push
## docker-push: pushes the docker image to registry
docker-push: check-environment docker-build
	docker push ${REGISTRY}/${ENV}/${APP}:${COMMIT_SHA}

.PHONY: help
## help: Prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
