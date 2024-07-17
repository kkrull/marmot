package svcfs_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kkrull/marmot/svcfs"
	expect "github.com/kkrull/marmot/testsupportexpect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JsonMetaRepoAdmin", func() {
	var (
		subject      *svcfs.JsonMetaRepoAdmin
		metaRepoPath string
		testFsRoot   string
	)

	BeforeEach(func() {
		testFsRoot = expect.NoError(os.MkdirTemp("", "JsonMetaDataRepo-"))
		metaRepoPath = filepath.Join(testFsRoot, "meta")
		DeferCleanup(os.RemoveAll, testFsRoot)
	})

	Describe("#Create", func() {
		It("returns an error, given a path that already exists", func() {
			Expect(os.Create(metaRepoPath)).NotTo(BeNil())

			subject = svcfs.NewJsonMetaRepoAdmin()
			Expect(subject.Create(metaRepoPath)).To(
				MatchError(fmt.Sprintf("path already exists: %s", metaRepoPath)))
		})

		It("returns an error when unable to check if the path exists", func() {
			subject = svcfs.NewJsonMetaRepoAdmin()
			invalidPathErr := subject.Create("\000x")
			Expect(invalidPathErr).NotTo(BeNil())
		})

		It("returns an error when creating files fails", func() {
			Expect(os.Chmod(testFsRoot, 0o555)).To(Succeed())

			subject = svcfs.NewJsonMetaRepoAdmin()
			Expect(subject.Create(metaRepoPath)).To(
				MatchError(ContainSubstring(fmt.Sprintf("failed to make directory %s", metaRepoPath))))
		})

		It("creates files in the meta repository and returns nil, otherwise", func() {
			subject = svcfs.NewJsonMetaRepoAdmin()
			Expect(subject.Create(metaRepoPath)).To(Succeed())

			metaDataDir := filepath.Join(metaRepoPath, ".marmot")
			Expect(os.Stat(metaDataDir)).NotTo(BeNil())

			metaDataFile := filepath.Join(metaDataDir, "meta-repo.json")
			Expect(os.Stat(metaDataFile)).NotTo(BeNil())
		})
	})
})
