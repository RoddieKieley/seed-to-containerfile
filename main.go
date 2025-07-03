package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func main() {
    dataDir := "data"
    if len(os.Args) > 1 {
        dataDir = os.Args[1]
    }

    entries, err := os.ReadDir(dataDir)
    if err != nil {
        log.Fatalf("failed to read data directory %s: %v", dataDir, err)
    }

    var failures int
    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }
        ext := filepath.Ext(entry.Name())
        if ext != ".yaml" && ext != ".yml" {
            continue
        }

        filePath := filepath.Join(dataDir, entry.Name())
        if _, err := ParseSpecData(filePath); err != nil {
            failures++
            fmt.Printf("FAIL  %s: %v\n", filePath, err)
        } else {
            fmt.Printf("SUCCESS %s\n", filePath)
        }
    }

    if failures > 0 {
        os.Exit(1)
    }
} 