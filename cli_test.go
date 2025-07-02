package main

import (
    "encoding/json"
    "os"
    "path/filepath"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

var _ = Describe("ParseSeedData", func() {
    It("parses example-data.json without error and captures ID", func() {
        targetPath := filepath.Join("data", "example-data.json")
        seed, err := ParseSeedData(targetPath)
        Expect(err).ToNot(HaveOccurred())
        Expect(seed).ToNot(BeNil())
        Expect(seed.ID).ToNot(BeNil())
        Expect(*seed.ID).To(Equal("01129bff-3d65-4e3d-8e82-6f2f269f818c"))
    })

    It("rewrites example-data.json and produces semantically identical JSON", func() {
        originalPath := filepath.Join("data", "example-data.json")
        rewrittenPath := filepath.Join("data", "example-data-rewritten.json")

        // Parse original
        seed, err := ParseSeedData(originalPath)
        Expect(err).ToNot(HaveOccurred())

        // Marshal back to JSON (pretty format)
        buf, err := json.MarshalIndent(seed, "", "  ")
        Expect(err).ToNot(HaveOccurred())

        // Write rewritten file
        Expect(os.WriteFile(rewrittenPath, buf, 0o644)).To(Succeed())

        // Load both files into generic structures for deep-equality comparison
        var originalObj, rewrittenObj interface{}

        originalBytes, err := os.ReadFile(originalPath)
        Expect(err).ToNot(HaveOccurred())
        Expect(json.Unmarshal(originalBytes, &originalObj)).To(Succeed())
        Expect(json.Unmarshal(buf, &rewrittenObj)).To(Succeed())

        Expect(rewrittenObj).To(Equal(originalObj))

        // Clean up rewritten file
        Expect(os.Remove(rewrittenPath)).To(Succeed())
    })
}) 