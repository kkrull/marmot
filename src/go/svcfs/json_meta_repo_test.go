package svcfs_test

import (
	"errors"
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
		subject    *svcfs.JsonMetaRepo
		testFsRoot string
	)

	createMetaRepo := func(path string) error {
		admin := jsonMetaRepoAdmin(nil)
		return admin.Create(path)
	}

	validPath := func() []string {
		return []string{os.TempDir()}
	}

	BeforeEach(func() {
		testFsRoot = expect.NoError(os.MkdirTemp("", "JsonMetaDataRepo-"))
		DeferCleanup(os.RemoveAll, testFsRoot)
	})

	Describe("#AddLocals", func() {
		It("skips duplicate paths, given 2 or more of the same exact path", func() {
			Expect(createMetaRepo(testFsRoot)).To(Succeed())
			subject = svcfs.NewJsonMetaRepo(testFsRoot)
			subject.AddLocals([]string{
				"/home/me/git/duplicate",
				"/home/me/git/other",
				"/home/me/git/duplicate",
			})

			listing := expect.NoError(subject.ListLocal())
			Expect(listing.LocalPaths()).To(ConsistOf(
				"/home/me/git/duplicate",
				"/home/me/git/other",
			))
		})

		It("returns an error, given a meta repo path that does not exist", func() {
			missingPath := filepath.Join(testFsRoot, "missing")
			_, statErr := os.Stat(missingPath)
			Expect(errors.Is(statErr, os.ErrNotExist)).To(BeTrue())

			subject = svcfs.NewJsonMetaRepo(missingPath)
			Expect(subject.AddLocals(validPath())).To(
				MatchError(ContainSubstring("failed to read file")))
		})
	})

	Describe("#AddRemotes", func() {
		BeforeEach(func() {
			Expect(createMetaRepo(testFsRoot)).To(Succeed())
			subject = svcfs.NewJsonMetaRepo(testFsRoot)
		})

		It("skips duplicate URLs, given 2 or more of the same URL", func() {
			subject.AddRemotes(testdata.NewURLs(
				"https://github.com/me/duplicate",
				"https://github.com/me/duplicate",
				"https://github.com/me/distinct",
			))

			listing := expect.NoError(subject.ListRemote())
			Expect(listing.RemoteHrefs()).To(ConsistOf(
				"https://github.com/me/duplicate",
				"https://github.com/me/distinct",
			))
		})
	})

	Context("when no repositories have been registered", func() {
		BeforeEach(func() {
			Expect(createMetaRepo(testFsRoot)).To(Succeed())
			subject = svcfs.NewJsonMetaRepo(testFsRoot)
		})

		It("#ListLocal returns empty", func() {
			repositories := expect.NoError(subject.ListLocal())
			Expect(repositories.Count()).To(Equal(0))
		})

		It("#ListRemote returns empty", func() {
			repositories := expect.NoError(subject.ListRemote())
			Expect(repositories.Count()).To(Equal(0))
		})
	})

	Context("when local repositories have been registered", func() {
		BeforeEach(func() {
			Expect(createMetaRepo(testFsRoot)).To(Succeed())
			subject = svcfs.NewJsonMetaRepo(testFsRoot)
			Expect(subject.AddLocals([]string{"/path/to/one"})).To(Succeed())
		})

		It("#AddLocals skips paths that have already been added", func() {
			subject.AddLocals([]string{
				"/path/to/one",
				"/path/to/two",
			})

			listing := expect.NoError(subject.ListLocal())
			Expect(listing.LocalPaths()).To(ConsistOf(
				"/path/to/one",
				"/path/to/two",
			))
		})

		It("#ListLocal includes each registered local repository", func() {
			returned := expect.NoError(subject.ListLocal())
			Expect(returned.LocalPaths()).To(ConsistOf("/path/to/one"))
		})
	})

	Context("when remote repositories have been registered", func() {
		BeforeEach(func() {
			Expect(createMetaRepo(testFsRoot)).To(Succeed())
			subject = svcfs.NewJsonMetaRepo(testFsRoot)

			given := testdata.NewURLs("https://github.com/me/a", "https://github.com/me/b")
			Expect(subject.AddRemotes(given)).To(Succeed())
		})

		It("#AddRemotes skips URLs that have already been added", func() {
			subject.AddRemotes(testdata.NewURLs("https://github.com/me/a"))

			listing := expect.NoError(subject.ListRemote())
			Expect(listing.RemoteHrefs()).To(ConsistOf(
				"https://github.com/me/a",
				"https://github.com/me/b",
			))
		})

		It("#ListRemote includes each registered remote", func() {
			listing := expect.NoError(subject.ListRemote())
			Expect(listing.RemoteHrefs()).To(ConsistOf(
				"https://github.com/me/a",
				"https://github.com/me/b",
			))
		})
	})
})
