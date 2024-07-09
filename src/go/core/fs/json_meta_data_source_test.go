package fs_test

import (
	"github.com/kkrull/marmot-core/fs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JsonMetaDataSource", func() {
	Describe("#Init", func() {
		PIt("returns an error if the directory already exists")

		PIt("returns an error when creating files fails")

		It("returns nil, otherwise", func() {
			subject := fs.JsonMetaDataSource{}
			Expect(subject.Init()).To(BeNil())
		})
	})
})
