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

// Package entry.
type Package struct {
    RegistryName string `json:"registry_name" yaml:"registry_name"`
    Name         string `json:"name" yaml:"name"`
    Version      string `json:"version" yaml:"version"`
    // Additional fields may exist but are omitted for brevity
}

// Remote describes remote endpoint information.
type Remote struct {
    Name string `json:"name,omitempty" yaml:"name,omitempty"`
    URL  string `json:"url,omitempty" yaml:"url,omitempty"`
} 