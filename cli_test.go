package main

import (
    "gopkg.in/yaml.v3"
    "os"
    "path/filepath"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "seed-to-containerfile/types"
)

var _ = Describe("ParseSeedData", func() {
    It("parses example-data.yaml without error and captures ID", func() {
        targetPath := filepath.Join("data", "example-data.yaml")
        spec, err := ParseSpecData(targetPath)
        Expect(err).ToNot(HaveOccurred())
        Expect(spec).ToNot(BeNil())
        Expect(spec.ServerDetail.ID).To(Equal("0007544a-3948-4934-866b-b4a96fe53b55"))
    })

    It("rewrites example-data.yaml and produces semantically identical YAML", func() {
        originalPath := filepath.Join("data", "example-data.yaml")
        rewrittenPath := filepath.Join("data", "example-data-rewritten.yaml")

        // Parse original
        spec, err := ParseSpecData(originalPath)
        Expect(err).ToNot(HaveOccurred())

        // Marshal back to YAML (pretty format)
        buf, err := yaml.Marshal(&struct{Spec *types.McpServerSpec `yaml:"spec"`}{Spec: spec})
        Expect(err).ToNot(HaveOccurred())

        // Write rewritten file
        Expect(os.WriteFile(rewrittenPath, buf, 0o644)).To(Succeed())

        // Load both files into generic structures for deep-equality comparison
        var originalObj, rewrittenObj interface{}

        originalBytes, err := os.ReadFile(originalPath)
        Expect(err).ToNot(HaveOccurred())
        Expect(yaml.Unmarshal(originalBytes, &originalObj)).To(Succeed())
        Expect(yaml.Unmarshal(buf, &rewrittenObj)).To(Succeed())

        Expect(rewrittenObj).To(Equal(originalObj))

        // Clean up rewritten file
        Expect(os.Remove(rewrittenPath)).To(Succeed())
    })
}) 