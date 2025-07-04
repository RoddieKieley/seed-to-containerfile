package types

// McpServerSpec defines the desired state of an MCP Server.
// Mirrors structure from mcp-catalog-operator but trimmed for this use-case.
// +kubebuilder:object:generate=true
// +kubebuilder:validation:Optional
// +kubebuilder:resource:path=mcpservers,scope=Namespaced
// +kubebuilder:subresource:status
// NOTE: json tags are used for YAML marshalling/unmarshalling.
type McpServerSpec struct {
    ServerDetail ServerDetail `json:"server_detail" yaml:"server_detail"`
}

// ServerDetail represents detailed server information.
type ServerDetail struct {
    Server   `json:",inline" yaml:",inline"`
    Packages []Package `json:"packages,omitempty" yaml:"packages,omitempty"`
    Remotes  []Remote  `json:"remotes,omitempty" yaml:"remotes,omitempty"`
}

// Server contains basic identifying metadata for a server.
type Server struct {
    Description string `json:"description,omitempty" yaml:"description,omitempty"`
    ID          string `json:"id,omitempty" yaml:"id,omitempty"`
    Name        string `json:"name,omitempty" yaml:"name,omitempty"`
    Repository  Repository `json:"repository,omitempty" yaml:"repository,omitempty"`
    VersionDetail VersionDetail `json:"version_detail,omitempty" yaml:"version_detail,omitempty"`
}

// Repository provides SCM repository details.
type Repository struct {
    URL    string `json:"url,omitempty" yaml:"url,omitempty"`
    Source string `json:"source,omitempty" yaml:"source,omitempty"`
    ID     string `json:"id,omitempty" yaml:"id,omitempty"`
}

// VersionDetail describes version information about the server.
type VersionDetail struct {
    Version     string `json:"version,omitempty" yaml:"version,omitempty"`
    ReleaseDate string `json:"release_date,omitempty" yaml:"release_date,omitempty"`
    IsLatest    bool   `json:"is_latest,omitempty" yaml:"is_latest,omitempty"`
}

// UserInput represents a user input as defined in the CRD schema.
type UserInput struct {
    // Core string fields
    Name        string            `json:"name,omitempty" yaml:"name,omitempty"`
    Description string            `json:"description,omitempty" yaml:"description,omitempty"`
    Value       string            `json:"value,omitempty" yaml:"value,omitempty"`
    Default     string            `json:"default,omitempty" yaml:"default,omitempty"`
    Format      string            `json:"format,omitempty" yaml:"format,omitempty"`

    // Flags
    IsRequired bool `json:"is_required,omitempty" yaml:"is_required,omitempty"`
    IsSecret   bool `json:"is_secret,omitempty" yaml:"is_secret,omitempty"`

    // Advanced
    Choices    []string          `json:"choices,omitempty" yaml:"choices,omitempty"`
    Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
    Template   string            `json:"template,omitempty" yaml:"template,omitempty"`

    // Nested variable substitutions (string->UserInput)
    Variables map[string]UserInput `json:"variables,omitempty" yaml:"variables,omitempty"`
}

// EnvironmentVariable represents an environment variable definition within a Package.
// It re-uses the UserInput schema but is kept separate for clarity in case of future divergence.
type EnvironmentVariable UserInput

// RuntimeArgument defines an argument (positional or named) for package/runtime execution.
type RuntimeArgument struct {
    // Required discriminator
    Type string `json:"type" yaml:"type"`

    // Common to UserInput
    Name        string            `json:"name,omitempty" yaml:"name,omitempty"`
    Description string            `json:"description,omitempty" yaml:"description,omitempty"`
    Value       string            `json:"value,omitempty" yaml:"value,omitempty"`
    Default     string            `json:"default,omitempty" yaml:"default,omitempty"`
    Format      string            `json:"format,omitempty" yaml:"format,omitempty"`

    // Flags
    IsRequired bool `json:"is_required,omitempty" yaml:"is_required,omitempty"`
    IsSecret   bool `json:"is_secret,omitempty" yaml:"is_secret,omitempty"`
    IsRepeated bool `json:"is_repeated,omitempty" yaml:"is_repeated,omitempty"`

    // Advanced
    Choices    []string          `json:"choices,omitempty" yaml:"choices,omitempty"`
    ValueHint  string            `json:"value_hint,omitempty" yaml:"value_hint,omitempty"`
    Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
    Template   string            `json:"template,omitempty" yaml:"template,omitempty"`

    Variables map[string]UserInput `json:"variables,omitempty" yaml:"variables,omitempty"`
}

// Package entry.
type Package struct {
    RegistryName string `json:"registry_name" yaml:"registry_name"`
    Name         string `json:"name" yaml:"name"`
    Version      string `json:"version" yaml:"version"`

    // Extended fields
    EnvironmentVariables []EnvironmentVariable `json:"environment_variables,omitempty" yaml:"environment_variables,omitempty"`
    PackageArguments     []RuntimeArgument     `json:"package_arguments,omitempty" yaml:"package_arguments,omitempty"`
    RuntimeArguments     []RuntimeArgument     `json:"runtime_arguments,omitempty" yaml:"runtime_arguments,omitempty"`
    RuntimeHint          string               `json:"runtime_hint,omitempty" yaml:"runtime_hint,omitempty"`
}

// Remote describes a remote connection endpoint.
type Remote struct {
    TransportType string       `json:"transport_type" yaml:"transport_type"`
    URL           string       `json:"url" yaml:"url"`
    Headers       []UserInput  `json:"headers,omitempty" yaml:"headers,omitempty"`
} 