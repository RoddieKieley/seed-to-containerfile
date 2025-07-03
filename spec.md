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

* The data provided is in the YAML 1.2 data language as per the definition of the YAML specification v1.2.2 at https://yaml.org/spec/1.2.2/
* Example provided YAML data is as follows:
```
spec:
  server_detail:
    description: Awesome MCP Servers - A curated list of Model Context Protocol servers
    id: 0007544a-3948-4934-866b-b4a96fe53b55
    name: io.github.appcypher/awesome-mcp-servers
    packages:
      - name: appcypher/awesome-mcp-servers
        registry_name: unknown
        version: ''
    repository:
      id: '895801050'
      source: github
      url: 'https://github.com/appcypher/awesome-mcp-servers'
    version_detail:
      is_latest: true
      release_date: '2025-05-16T19:16:40Z'
      version: 0.0.1-seed
```

# What to do

## Create the command line utility

### Process the data into golang types

* Write the Example provided YAML data to a file `example-data.yaml` in the `data` directory within the project. If this file is found to already exist in the project then overwrite the existing `example-data.yaml` file with the Example provided YAML data previously provided.
* Analyse the Example provided YAML data file `example-data.yaml`, creating matching valid golang structs with json: annotations.
** The top level `spec:` entry corresponds to a specification for an MCP Server. The golang struct utilized in this command line utility project should match the `type McpServerSpec struct` found at https://github.com/RHEcosystemAppEng/mcp-catalog-operator/blob/main/api/v1alpha1/mcpserver_types.go.
** Further more any subsequent related golang struct types that are required should also be used directly from https://github.com/RHEcosystemAppEng/mcp-catalog-operator/blob/main/api/v1alpha1/mcpserver_types.go where possible.
** Write these required golang struct types to a file `mcpserver_types.go` in the `types` directory.
** Example golang struct type with json: annotation for the top level `spec:` and child `server_detail:` sections from the Example provided YAML data previously provided:
```
// McpServerSpec defines the desired state of McpServer.
type McpServerSpec struct {
	ServerDetail ServerDetail `json:"server_detail"`
}

// ServerDetail represents detailed server information as defined in the spec
type ServerDetail struct {
	Server   `json:",inline"`
	Packages []Package `json:"packages,omitempty"`
	Remotes  []Remote  `json:"remotes,omitempty"`
}
```
** Where the golang struct types differ between the Example golang struct type and those in the https://github.com/RHEcosystemAppEng/mcp-catalog-operator/blob/main/api/v1alpha1/mcpserver_types.go file differ, prefer those in the mcpserver_types.go file where possible.


### Write the golang source code for the command line utility

* Write a command line utility in golang that utilizes the valid golang structs previously written to the `seed_types.go` file in the `types` directory.
** This command line utility in golang must:
*** Examine the `data` directory to find all `*.yaml` files.
*** For each of the `*.yaml` files found in the `data` directory, read and parse the YAML data from the individual file into the valid golang structs found in the `seed_types.go` file.
**** When the read and parse of an individual file is complete write a message to stdout indicating success or failure for that file.

### Write the golang test source code for the command line utility

* You prefer to write the golang test source code for the command line utility using the ginkgo Modern Testing Framework from https://github.com/onsi/ginkgo and the gomega Preferred Matcher Library from https://github.com/onsi/gomega. The ginkgo Modern Testing Framework focuses on Behaviour-driven development as documented at https://en.wikipedia.org/wiki/Behavior-driven_development
* Write matching tests for this golang command line utility using the ginkgo Modern Testing Framework and gomega Preferred Matcher Library using a Behaviour-driven development approach.
** In particular ensure the tests cover the functionality reading, parsing, and marshalling the YAML into golang struct type instances from `seed_types.go` by using the original Example provided YAML data from `example-data.yaml` in the `data` directory.
** A test must be provided that takes the Example provided YAML data from `example-data.yaml` in the `data` directory, loads it successfully into the golang struct types in `seed_types.go`, then writes the data back to a new file named `example-data-rewritten.yaml` in the `data` directory in correct YAML format such that it is identical and compares successfully against the `example-data.yaml` in the `data` directory.
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
