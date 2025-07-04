package main

import (
    "gopkg.in/yaml.v3"
    "os"
    "path/filepath"
    "strings"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "seed-to-containerfile/types"
)

var _ = Describe("ParseSeedData", func() {
    It("parses and rewrites each example-data-#.yaml without loss", func() {
        files, err := filepath.Glob(filepath.Join("data", "example-data-*.yaml"))
        Expect(err).ToNot(HaveOccurred())
        Expect(files).ToNot(BeEmpty())

        for _, originalPath := range files {
            By("processing " + originalPath)

            spec, err := ParseSpecData(originalPath)
            Expect(err).ToNot(HaveOccurred())

            // Marshal back to YAML
            buf, err := yaml.Marshal(&struct{Spec *types.McpServerSpec `yaml:"spec"`}{Spec: spec})
            Expect(err).ToNot(HaveOccurred())

            rewrittenPath := strings.Replace(originalPath, "example-data-", "example-data-rewritten-", 1)
            Expect(os.WriteFile(rewrittenPath, buf, 0o644)).To(Succeed())

            rewrittenSpec, err := ParseSpecData(rewrittenPath)
            Expect(err).ToNot(HaveOccurred())
            Expect(rewrittenSpec).To(Equal(spec))

            Expect(os.Remove(rewrittenPath)).To(Succeed())
        }
    })
}) 