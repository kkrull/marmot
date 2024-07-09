package fs_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JsonMetaDataSource", func() {
	Describe("#Init", func() {
		PIt("returns an error if the directory already exists")
		PIt("returns an error when creating files fails")
		PIt("returns nil, after creating files in a directory that does not exist")
	})
})
