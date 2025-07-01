package main

import (
    "encoding/json"
    "os"
    "seed-to-containerfile/types"
)

// ParseSeedData reads the JSON seed file at the supplied path and unmarshals it
// into a SeedData structure.
func ParseSeedData(filePath string) (*types.SeedData, error) {
    raw, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }
    var seed types.SeedData
    if err := json.Unmarshal(raw, &seed); err != nil {
        return nil, err
    }
    return &seed, nil
} 