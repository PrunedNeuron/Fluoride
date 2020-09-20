ARG TARGETOS=linux
ARG TARGETARCH=amd64

FROM golang:alpine AS builder
RUN apk update && apk upgrade && \
	apk add bash git make coreutils

WORKDIR /go/src/github.com/PrunedNeuron/Fluoride/
COPY . .

RUN make

# Build minimal image, chose busybox because it has /bin/sh and is way
# smaller than alpine
FROM busybox
WORKDIR /go/src/github.com/PrunedNeuron/Fluoride/

COPY --from=builder /go/src/github.com/PrunedNeuron/Fluoride/bin ./bin
COPY --from=builder /go/src/github.com/PrunedNeuron/Fluoride/scripts ./scripts
COPY --from=builder /go/src/github.com/PrunedNeuron/Fluoride/config.yml .

ENTRYPOINT [ "scripts/docker_entrypoint.sh" ]
CMD [ "serve" ]
