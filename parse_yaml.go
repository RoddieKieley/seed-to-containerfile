package main

import (
    "os"

    "gopkg.in/yaml.v3"
    "seed-to-containerfile/types"
)

// ParseSpecData reads a YAML file and unmarshals it into McpServerSpec.
func ParseSpecData(filePath string) (*types.McpServerSpec, error) {
    raw, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }
    var wrapper struct{
        Spec types.McpServerSpec `yaml:"spec"`
    }
    if err := yaml.Unmarshal(raw, &wrapper); err != nil {
        return nil, err
    }
    return &wrapper.Spec, nil
} 