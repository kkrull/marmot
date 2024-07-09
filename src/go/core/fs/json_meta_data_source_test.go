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
			var fixtureErr error
			testFsRoot, fixtureErr = os.MkdirTemp("", "JsonMetaDataSource-")
			Expect(fixtureErr).To(BeNil())

			metaDataPath = filepath.Join(testFsRoot, "meta")
		})

		AfterEach(func() {
			if err := os.RemoveAll(testFsRoot); err != nil {
				fmt.Printf("JsonMetaDataSource test: failed to remove %s\n", testFsRoot)
				fmt.Println(err.Error())
			}
		})

		It("returns an error, given a path that already exists", func() {
			Expect(os.Create(metaDataPath)).NotTo(BeNil())

			subject := fs.JsonMetaDataSource{Path: metaDataPath}
			Expect(subject.Init()).To(
				MatchError(fmt.Sprintf("%s: path already exists", metaDataPath)))
		})

		It("returns an error when unable to check if the path exists", func() {
			subject := fs.JsonMetaDataSource{Path: "\000x"}
			invalidPathErr := subject.Init()
			Expect(invalidPathErr).NotTo(BeNil())
		})

		It("returns an error when creating files fails", Focus, func() {
			Expect(os.Chmod(testFsRoot, 0o555)).To(BeNil())

			subject := fs.JsonMetaDataSource{Path: metaDataPath}
			Expect(subject.Init()).To(
				MatchError(ContainSubstring(fmt.Sprintf("createMetaData %s", metaDataPath))))
		})

		It("creates a meta repository and returns nil, otherwise", func() {
			subject := fs.JsonMetaDataSource{Path: metaDataPath}
			Expect(subject.Init()).To(BeNil())

			stat, statErr := os.Stat(metaDataPath)
			Expect(statErr).To(BeNil())
			Expect(stat).NotTo(BeNil())
		})
	})
})
