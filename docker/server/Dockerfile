ARG TARGETOS=linux
ARG TARGETARCH=amd64

FROM golang:alpine AS builder

RUN apk update && apk upgrade && \
	apk add bash git make coreutils

WORKDIR /go/src/github.com/PrunedNeuron/Fluoride/

# COPY the source code
COPY . .

# Download and verify go dependencies
RUN go mod download
RUN go mod verify

# Build binary
RUN make

# Build minimal image
FROM scratch
WORKDIR /go/src/github.com/PrunedNeuron/Fluoride/

# Retrieve the binary from the previous stage
COPY --from=builder /go/src/github.com/PrunedNeuron/Fluoride/bin ./bin
COPY --from=builder /go/src/github.com/PrunedNeuron/Fluoride/config.yml .

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "./bin/fluoride", "serve" ]