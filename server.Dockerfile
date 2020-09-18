FROM golang:alpine AS builder

# Add all the source code (except what's ignored
# under `.dockerignore`) to the build context.
ADD ./ /go/src/github.com/PrunedNeuron/Fluoride/

RUN set -ex && \
	cd /go/src/github.com/PrunedNeuron/Fluoride && \       
	CGO_ENABLED=0 go build \
	-tags netgo \
	# -v -a \
	-v \
	-ldflags '-extldflags "-static"' && \
	mv ./Fluoride /usr/bin/Fluoride

FROM alpine:latest
RUN apk add --update ca-certificates

# Retrieve the binary from the previous stage
COPY --from=builder /usr/bin/Fluoride /usr/local/bin/Fluoride

EXPOSE 3000
# Set the binary as the entrypoint of the container
ENTRYPOINT [ "Fluoride", "serve" ]