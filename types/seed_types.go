package types

// SeedData represents the root structure of the JSON seed file.
type SeedData struct {
    ID           *string            `json:"id,omitempty"`
    Name         *string            `json:"name,omitempty"`
    Description  *string            `json:"description,omitempty"`
    Repository   *RepositoryType    `json:"repository,omitempty"`
    VersionDetail *VersionDetailType `json:"version_detail,omitempty"`
    Packages     []PackageType      `json:"packages,omitempty"`
}

// RepositoryType captures repository information.
type RepositoryType struct {
    URL    *string `json:"url,omitempty"`
    Source *string `json:"source,omitempty"`
    ID     *string `json:"id,omitempty"`
}

// VersionDetailType captures version information.
type VersionDetailType struct {
    Version     *string `json:"version,omitempty"`
    ReleaseDate *string `json:"release_date,omitempty"`
    IsLatest    *bool   `json:"is_latest,omitempty"`
}

// PackageType describes an individual package.
type PackageType struct {
    RegistryName    *string              `json:"registry_name,omitempty"`
    Name            *string              `json:"name,omitempty"`
    Version         *string              `json:"version,omitempty"`
    PackageArguments []PackageArgumentType `json:"package_arguments,omitempty"`
}

// PackageArgumentType details a single argument for a package.
type PackageArgumentType struct {
    Description *string `json:"description,omitempty"`
    IsRequired  *bool   `json:"is_required,omitempty"`
    Format      *string `json:"format,omitempty"`
    Value       *string `json:"value,omitempty"`
    Default     *string `json:"default,omitempty"`
    ArgType     *string `json:"type,omitempty"`
    ValueHint   *string `json:"value_hint,omitempty"`
} 