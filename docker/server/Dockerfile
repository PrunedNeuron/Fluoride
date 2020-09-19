FROM golang:alpine AS builder

# Add all the source code (except what's ignored
# under `.dockerignore`) to the build context.

WORKDIR /go/src/github.com/PrunedNeuron/Fluoride/

# Set args 
# RUN set -ex

COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download
# COPY the source code
COPY . .

ARG TARGETOS
ARG TARGETARCH

# Build binary
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -tags netgo -ldflags '-extldflags "-static"'

FROM alpine:latest
RUN apk add --update ca-certificates

WORKDIR /go/src/github.com/PrunedNeuron/Fluoride/

# Retrieve the binary from the previous stage
COPY --from=builder /go/src/github.com/PrunedNeuron/Fluoride/ .

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "./Fluoride", "serve" ]