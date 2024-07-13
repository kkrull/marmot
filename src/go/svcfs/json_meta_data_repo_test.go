package svcfs_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kkrull/marmot/svcfs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JsonMetaDataRepo", func() {
	Describe("#Init", func() {
		var (
			metaRepoPath string
			testFsRoot   string
		)

		BeforeEach(func() {
			var fixtureErr error
			testFsRoot, fixtureErr = os.MkdirTemp("", "JsonMetaDataRepo-")
			Expect(fixtureErr).To(BeNil())

			metaRepoPath = filepath.Join(testFsRoot, "meta")
		})

		AfterEach(func() {
			if err := os.RemoveAll(testFsRoot); err != nil {
				fmt.Printf("JsonMetaDataRepo test: failed to remove %s\n", testFsRoot)
				fmt.Println(err.Error())
			}
		})

		It("returns an error, given a path that already exists", func() {
			Expect(os.Create(metaRepoPath)).NotTo(BeNil())

			subject := svcfs.JsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init()).To(
				MatchError(fmt.Sprintf("%s: path already exists", metaRepoPath)))
		})

		It("returns an error when unable to check if the path exists", func() {
			subject := svcfs.JsonMetaDataRepo("\000x")
			invalidPathErr := subject.Init()
			Expect(invalidPathErr).NotTo(BeNil())
		})

		It("returns an error when creating files fails", func() {
			Expect(os.Chmod(testFsRoot, 0o555)).To(Succeed())

			subject := svcfs.JsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init()).To(
				MatchError(ContainSubstring(fmt.Sprintf("createMetaData %s", metaRepoPath))))
		})

		It("creates a meta repository and returns nil, otherwise", func() {
			subject := svcfs.JsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init()).To(Succeed())

			metaDataDir := filepath.Join(metaRepoPath, ".marmot")
			Expect(os.Stat(metaDataDir)).NotTo(BeNil())

			metaDataFile := filepath.Join(metaDataDir, "meta-repo.json")
			Expect(os.Stat(metaDataFile)).NotTo(BeNil())
		})
	})
})
