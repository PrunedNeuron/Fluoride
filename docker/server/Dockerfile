ARG TARGETOS=linux
ARG TARGETARCH=amd64

FROM golang:alpine AS builder
RUN apk update && apk upgrade && \
	apk add make

WORKDIR /go/src/github.com/PrunedNeuron/Fluoride/

COPY go.mod .
COPY go.mod .

RUN go mod download

COPY . .

RUN make

# Build minimal image
# Busybox has /bin/sh and is smaller than alpine
FROM busybox
WORKDIR /go/src/github.com/PrunedNeuron/Fluoride/

COPY --from=builder /go/src/github.com/PrunedNeuron/Fluoride/bin ./bin
COPY --from=builder /go/src/github.com/PrunedNeuron/Fluoride/scripts ./scripts
COPY --from=builder /go/src/github.com/PrunedNeuron/Fluoride/config.yml .

ENTRYPOINT [ "scripts/docker_entrypoint.sh" ]
CMD [ "serve" ]
