package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)

func main() {
    filePath := "data/example-data.json"
    if len(os.Args) > 1 {
        filePath = os.Args[1]
    }

    seed, err := ParseSeedData(filePath)
    if err != nil {
        log.Fatalf("failed to parse seed data: %v", err)
    }

    pretty, err := json.MarshalIndent(seed, "", "  ")
    if err != nil {
        log.Fatalf("failed to marshal seed data: %v", err)
    }
    fmt.Println(string(pretty))
} 