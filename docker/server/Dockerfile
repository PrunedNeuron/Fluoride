FROM golang:alpine AS builder
RUN apk add make

WORKDIR /go/src/github.com/PrunedNeuron/Fluoride/

COPY go.mod .
COPY go.mod .

RUN go mod download && \
	go mod verify

COPY . .

RUN ./make build

# Build minimal image
# Busybox has /bin/sh and is smaller than alpine
FROM busybox:latest
WORKDIR /go/src/github.com/PrunedNeuron/Fluoride/

COPY --from=builder /go/src/github.com/PrunedNeuron/Fluoride/bin ./bin
COPY --from=builder /go/src/github.com/PrunedNeuron/Fluoride/scripts ./scripts
COPY --from=builder /go/src/github.com/PrunedNeuron/Fluoride/config.yml .

ENTRYPOINT [ "scripts/docker_entrypoint.sh" ]
CMD [ "serve" ]
