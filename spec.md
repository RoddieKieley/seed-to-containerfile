# Persona

* You are a Senior Software Engineer.
* You have expertise in the area of writing golang code and containerizing command line utilities.
* You prefer to utilize the buildah tool at https://github.com/containers/buildah to build OCI images.
* You prefer to utilize the podman tool at https://github.com/containers/podman to run OCI images.

# Rules

* Think, plan, and reason step by step.
* Do not jump to conclusions. 
* Once the plan is complete review each of the generated steps for accuracy. If a given step is found to be inaccurate then fix that step and any subsequent steps that transitively require updates.

# Data provided

* The data provided is in the JavaScript Object Notation (JSON) as per the Request for Comments: 4627 at https://www.ietf.org/rfc/rfc4627.txt
* Example provided JSON data is as follows:
```
{
  "id": "01129bff-3d65-4e3d-8e82-6f2f269f818c",
  "name": "io.github.gongrzhe/redis-mcp-server",
  "description": "A Redis MCP server (pushed to https://github.com/modelcontextprotocol/servers/tree/main/src/redis) implementation for interacting with Redis databases. This server enables LLMs to interact with Redis key-value stores through a set of standardized tools.",
  "repository": {
    "url": "https://github.com/GongRzhe/REDIS-MCP-Server",
    "source": "github",
    "id": "907849235"
  },
  "version_detail": {
    "version": "0.0.1-seed",
    "release_date": "2025-05-16T19:13:21Z",
    "is_latest": true
  },
  "packages": [
    {
      "registry_name": "docker",
      "name": "@gongrzhe/server-redis-mcp",
      "version": "1.0.0",
      "package_arguments": [
        {
          "description": "Docker image to run",
          "is_required": true,
          "format": "string",
          "value": "mcp/redis",
          "default": "mcp/redis",
          "type": "positional",
          "value_hint": "mcp/redis"
        },
        {
          "description": "Redis server connection string",
          "is_required": true,
          "format": "string",
          "value": "redis://host.docker.internal:6379",
          "default": "redis://host.docker.internal:6379",
          "type": "positional",
          "value_hint": "host.docker.internal:6379"
        }
      ]
    }
  ]
}
```

# What to do

## Create the command line utility

### Process the data into golang types

* Write the Example provided JSON data to a file `example-data.json` in the `data` directory within the project. If this file is found to already exist in the project then overwrite the existing `example-data.json` file with the Example provided JSON data previously provided.
* Analyse the Example provided JSON data file `example-data.json`, creating matching valid golang structs with json: annotations.
** Write these matching valid golang structs to a file `seed_types.go` in the `types` directory using the following Example as a guide for all the golang struct types required to read in the Example provided JSON data from the `example-data.json` file.
** Example golang struct with json: annotation for the "version_detail" entry from the Example provided JSON data previously provided:
```
type VersionDetailType struct {
    // "version": "0.0.1-seed",
    Version *string `json:"version,omitempty"`
    // "release_date": "2025-05-16T19:13:21Z",
    ReleaseDate *string `json:"releaseDate,omitempty"
    //"is_latest": true
    IsLatest *bool `json:"isLatest,omitempty"
}
```

### Write the golang source code for the command line utility

* Write a command line utility in golang that utilizes the valid golang structs previously written to the `seed_types.go` file in the `types` directory.
** This command line utility in golang must:
*** Examine the `data` directory to find all `*.json` files.
*** For each of the `*.json` files found in the `data` directory, read and parse the JSON data from the individual file into the valid golang structs found in the `seed_types.go` file.
**** When the read and parse of an individual file is complete write a message to stdout indicating success or failure for that file.

### Write the golang test source code for the command line utility

* You prefer to write the golang test source code for the command line utility using the ginkgo Modern Testing Framework from https://github.com/onsi/ginkgo and the gomega Preferred Matcher Library from https://github.com/onsi/gomega. The ginkgo Modern Testing Framework focuses on Behaviour-driven development as documented at https://en.wikipedia.org/wiki/Behavior-driven_development
* Write matching tests for this golang command line utility using the ginkgo Modern Testing Framework and gomega Preferred Matcher Library using a Behaviour-driven development approach.
** In particular ensure the tests cover the functionality reading, parsing, and marshalling the JSON into golang struct type instances from `seed_types.go` by using the original Example provided JSON data from `example-data.json` in the `data` directory.
** A test must be provided that takes the Example provided JSON data from `example-data.json` in the `data` directory, loads it successfully into the golang struct types in `seed_types.go`, then writes the data back to a new file named `example-data-rewritten.json` in the `data` directory in correct JSON format such that it is identical and compares successfully against the `example-data.json` in the `data` directory.
* Execute the tests created using the `go test ./...` shell command. If all tests are Passed then proceed to the `Build the OCI image` section.

## Build the OCI image

* This command line utility should build correctly via the `go build ./...` shell command. If not return to the `Create the command line utility` section.

### Create the Dockerfile

* The `Dockerfile` reference documentation is at https://docs.docker.com/reference/dockerfile/. The available instructions that can be used in the `Dockerfile` are listed in the Overview on that web page url.

* Write a `Dockerfile` to the root directory of the project. If it is found to exist, overwrite the existing `Dockerfile` in the root directory of the project.
** This `Dockerfile` will be a multi-stage build as described at https://docs.docker.com/get-started/docker-concepts/building-images/multi-stage-builds/ and at https://docs.docker.com/guides/golang/build-images/ for golang specifically.
** There are three stages in this multi-stage `Dockerfile`:
*** Build the command line utility from golang source in the first `build-stage`.
*** Run the tests in the container in the second `run-test-stage`.
*** Deploy the command line utility binary into a leaner image for later execution in the `build-release-stage`.
** When building a multi-stage golang based OCI image via a `Dockerfile`:
*** You prefer to use registry.access.redhat.com/ubi9/go-toolset:1.23.9 as the golang source code `build-stage` builder image.
*** You prefer to use registry.access.redhat.com/ubi9/ubi-micro:9.4 as the image to Deploy the command line utility binary into during the `build-release-stage`.
** During the `build-stage` the command line utility binary compiled via `go build` from the golang source of the project should be built in the `/app` directory.  

* Example multi-stage `Dockerfile` that builds from golang source, runs the tests in the container, and then copies the command line utility binary into a leaner image for later execution.
```
# syntax=docker/dockerfile:1

# Build the command line utility from golang source
FROM registry.access.redhat.com/ubi9/go-toolset:1.23.9 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the command line utility binary into a leaner image for later execution
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /docker-gs-ping /docker-gs-ping

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-ping"]
```
## For the build-stage in the `Dockerfile` utilize the registry.access.redhat.com/ubi9/go-toolset OCI image. Use the golang version that most closely matches that used

### Build the `Dockerfile` into the OCI image

* Use the instructions in Using Containerfiles/Dockerfiles with Buildah at https://github.com/containers/buildah/blob/main/docs/tutorials/01-intro.md#using-containerfilesdockerfiles-with-buildah to build the `Dockerfile` from the previous 'Create the Dockerfile` section.
** The name to be given to the created image is 'localhost/seedtocontainer' with tag `latest`

### Test the execution of the OCI image built from the `Dockerfile` using the following command:
```
podman run --rm localhost/seedtocontainer:latest
```
