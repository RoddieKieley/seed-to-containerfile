package main

import (
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
}) 