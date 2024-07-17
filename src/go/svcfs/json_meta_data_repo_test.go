package svcfs_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kkrull/marmot/svcfs"
	testdata "github.com/kkrull/marmot/testsupportdata"
	expect "github.com/kkrull/marmot/testsupportexpect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JsonMetaDataRepo", func() {
	var (
		subject      *svcfs.JsonMetaDataRepo
		metaRepoPath string
		testFsRoot   string
	)

	BeforeEach(func() {
		testFsRoot = expect.NoError(os.MkdirTemp("", "JsonMetaDataRepo-"))
		metaRepoPath = filepath.Join(testFsRoot, "meta")
		DeferCleanup(os.RemoveAll, testFsRoot)
	})

	Describe("#Init", func() {
		It("returns an error, given a path that already exists", func() {
			Expect(os.Create(metaRepoPath)).NotTo(BeNil())

			subject = svcfs.NewJsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init(metaRepoPath)).To(
				MatchError(fmt.Sprintf("path already exists: %s", metaRepoPath)))
		})

		It("returns an error when unable to check if the path exists", func() {
			subject = svcfs.NewJsonMetaDataRepo("\000x")
			invalidPathErr := subject.Init("\000x")
			Expect(invalidPathErr).NotTo(BeNil())
		})

		It("returns an error when creating files fails", func() {
			Expect(os.Chmod(testFsRoot, 0o555)).To(Succeed())

			subject = svcfs.NewJsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init(metaRepoPath)).To(
				MatchError(ContainSubstring(fmt.Sprintf("failed to make directory %s", metaRepoPath))))
		})

		It("creates a meta repository and returns nil, otherwise", func() {
			subject = svcfs.NewJsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init(metaRepoPath)).To(Succeed())

			metaDataDir := filepath.Join(metaRepoPath, ".marmot")
			Expect(os.Stat(metaDataDir)).NotTo(BeNil())

			metaDataFile := filepath.Join(metaDataDir, "meta-repo.json")
			Expect(os.Stat(metaDataFile)).NotTo(BeNil())
		})
	})

	Context("when no repositories have been registered", func() {
		BeforeEach(func() {
			subject = svcfs.NewJsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init(metaRepoPath)).To(Succeed())
		})

		It("#ListRemote returns empty", func() {
			repositories := expect.NoError(subject.ListRemote())
			Expect(repositories.Count()).To(Equal(0))
		})
	})

	Context("when remote repositories have been registered", func() {
		BeforeEach(func() {
			subject = svcfs.NewJsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init(metaRepoPath)).To(Succeed())

			Expect(subject.AddRemote(testdata.NewURL("https://github.com/me/a"))).To(Succeed())
			listOne := expect.NoError(subject.ListRemote())
			Expect(listOne.RemoteHrefs()).To(ConsistOf("https://github.com/me/a"))

			Expect(subject.AddRemote(testdata.NewURL("https://github.com/me/b"))).To(Succeed())
		})

		It("#ListRemote includes each registered remote", func() {
			listTwo := expect.NoError(subject.ListRemote())
			Expect(listTwo.RemoteHrefs()).To(ConsistOf(
				"https://github.com/me/a",
				"https://github.com/me/b",
			))
		})
	})
})
