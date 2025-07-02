# syntax=docker/dockerfile:1

############################
# Build the command line utility
############################
FROM registry.access.redhat.com/ubi9/go-toolset:1.23.9 AS build-stage

# Disable VCS stamping to avoid git requirement in container builds
ENV GOFLAGS="-buildvcs=false"

WORKDIR /app

# Copy go modules first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o /app/seedtocontainer ./

############################
# Run the unit tests inside container
############################
FROM build-stage AS run-test-stage
RUN go test -v ./...

############################
# Final lean runtime image
############################
FROM registry.access.redhat.com/ubi9/ubi-micro:9.4 AS build-release-stage

WORKDIR /

# Copy binary and data directory from build stage
COPY --from=build-stage /app/seedtocontainer /seedtocontainer
COPY --from=build-stage /app/data /data

# Run as non-root where possible
USER 1001

ENTRYPOINT ["/seedtocontainer"] 