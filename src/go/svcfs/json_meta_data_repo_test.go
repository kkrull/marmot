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
			Expect(subject.Init()).To(
				MatchError(fmt.Sprintf("path already exists: %s", metaRepoPath)))
		})

		It("returns an error when unable to check if the path exists", func() {
			subject = svcfs.NewJsonMetaDataRepo("\000x")
			invalidPathErr := subject.Init()
			Expect(invalidPathErr).NotTo(BeNil())
		})

		It("returns an error when creating files fails", func() {
			Expect(os.Chmod(testFsRoot, 0o555)).To(Succeed())

			subject = svcfs.NewJsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init()).To(
				MatchError(ContainSubstring(fmt.Sprintf("failed to make directory %s", metaRepoPath))))
		})

		It("creates a meta repository and returns nil, otherwise", func() {
			subject = svcfs.NewJsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init()).To(Succeed())

			metaDataDir := filepath.Join(metaRepoPath, ".marmot")
			Expect(os.Stat(metaDataDir)).NotTo(BeNil())

			metaDataFile := filepath.Join(metaDataDir, "meta-repo.json")
			Expect(os.Stat(metaDataFile)).NotTo(BeNil())
		})
	})

	Context("when no repositories have been registered", func() {
		BeforeEach(func() {
			subject = svcfs.NewJsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init()).To(Succeed())
		})

		It("#List returns empty", func() {
			repositories := expect.NoError(subject.List())
			Expect(repositories.Count()).To(Equal(0))
		})
	})

	Context("when remote repositories have been registered", func() {
		BeforeEach(func() {
			subject = svcfs.NewJsonMetaDataRepo(metaRepoPath)
			Expect(subject.Init()).To(Succeed())

			validUrl := testdata.NewURL("https://github.com/actions/checkout")
			Expect(subject.RegisterRemote(validUrl)).To(Succeed())
		})

		It("#List includes each registered remote", func() {
			listing := expect.NoError(subject.List())
			Expect(listing.RemoteHrefs()).To(ConsistOf("https://github.com/actions/checkout"))
		})

		It("can register and list two repositories", Pending, func() {
		})
	})
})
