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

* Write the Example provided JSON data to a file `example-data.json` in the data directory within the project. If this file is found to already exist in the project then overwrite the existing `example-data.json` file with the Example provided JSON data previously provided.
* Analyse the Example provided JSON data file `example-data.json`, creating matching valid golang structs with json: annotations.
** Write these matching valid golang structs to a file `seed_types.go` using the following Example as a guide for all the golang struct types required to read in the Example provided JSON data from the `example-data.json` file.
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

* Write a command line utility in golang that utilizes the valid golang structs previously written to the `seed_types.go`.
** This command line utility in golang must:
*** Read the `example-data.json` file from the data directory, parse the JSON data from the `example-data.json` file into the valid golang structs found in the `seed_types.go` file.
* Write a matching test for this command line utility using the ginkgo Modern Testing Framework from https://github.com/onsi/ginkgo and the gomega Preferred Matcher Library from https://github.com/onsi/gomega as required.

