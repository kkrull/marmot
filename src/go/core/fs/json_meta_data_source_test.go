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
			testFsRoot   string
			metaDataPath string
		)

		BeforeEach(func() {
			testFsRoot, fsErr := os.MkdirTemp("", "JsonMetaDataSource-")
			Expect(fsErr).To(BeNil())
			metaDataPath = filepath.Join(testFsRoot, "meta")
			// fmt.Printf("metaDataPath: %s\n", metaDataPath)
		})

		AfterEach(func() {
			if err := os.RemoveAll(testFsRoot); err != nil {
				fmt.Printf("JsonMetaDataSource test: failed to remove %s\n", testFsRoot)
				fmt.Println(err.Error())
			}
		})

		It("returns an error if the directory already exists", func() {
			// https://pkg.go.dev/os@go1.22.5#Mkdir
			// https://stackoverflow.com/questions/37932551/mkdir-if-not-exists-using-golang
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
