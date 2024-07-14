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

	Describe("#Init", func() {
		It("returns an error, given a path that already exists", func() {
			Expect(os.Create(metaRepoPath)).NotTo(BeNil())

			subject := svcfs.NewJsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init()).To(
				MatchError(fmt.Sprintf("path already exists: %s", metaRepoPath)))
		})

		It("returns an error when unable to check if the path exists", func() {
			subject := svcfs.NewJsonMetaDataRepo("\000x")
			invalidPathErr := subject.Init()
			Expect(invalidPathErr).NotTo(BeNil())
		})

		It("returns an error when creating files fails", func() {
			Expect(os.Chmod(testFsRoot, 0o555)).To(Succeed())

			subject := svcfs.NewJsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init()).To(
				MatchError(ContainSubstring(fmt.Sprintf("failed to make directory %s", metaRepoPath))))
		})

		It("creates a meta repository and returns nil, otherwise", func() {
			subject := svcfs.NewJsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init()).To(Succeed())

			metaDataDir := filepath.Join(metaRepoPath, ".marmot")
			Expect(os.Stat(metaDataDir)).NotTo(BeNil())

			metaDataFile := filepath.Join(metaDataDir, "meta-repo.json")
			Expect(os.Stat(metaDataFile)).NotTo(BeNil())
		})
	})

	Describe("#List", func() {
		It("returns empty repositories, given a meta repo in which none have been registered", func() {
			subject := svcfs.NewJsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init()).To(Succeed())

			if repositories, listErr := subject.List(); listErr != nil {
				Fail(listErr.Error())
			} else {
				Expect(repositories.Names()).To(BeEmpty())
			}
		})

		It("lists each known remote repository, given a meta repo with remote repositories", Pending, func() {
		})
	})
})
