package fs_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kkrull/marmot-core/fs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JsonMetaDataSource", func() {
	Describe("#Init", func() {
		var (
			metaDataPath string
			testFsRoot   string
		)

		BeforeEach(func() {
			testFsRoot, err := os.MkdirTemp("", "JsonMetaDataSource-")
			Expect(err).To(BeNil())
			metaDataPath = filepath.Join(testFsRoot, "meta")
		})

		AfterEach(func() {
			if err := os.RemoveAll(testFsRoot); err != nil {
				fmt.Printf("JsonMetaDataSource test: failed to remove %s\n", testFsRoot)
				fmt.Println(err.Error())
			}
		})

		It("returns an error when unable to check if the given path exists", func() {
			subject := fs.JsonMetaDataSource{Path: "\000x"}
			invalidPathErr := subject.Init()
			Expect(invalidPathErr).NotTo(BeNil())
		})

		It("returns an error, given a path that already exists", func() {
			Expect(os.MkdirAll(metaDataPath, os.ModePerm)).To(Succeed())

			subject := fs.JsonMetaDataSource{Path: metaDataPath}
			Expect(subject.Init()).To(
				MatchError(fmt.Sprintf("%s: path already exists", metaDataPath)))
		})

		PIt("returns an error when creating files fails")

		It("returns nil, otherwise", func() {
			subject := fs.JsonMetaDataSource{Path: metaDataPath}
			Expect(subject.Init()).To(BeNil())
		})
	})
})
