package svcfs_test

import (
	"os"

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

	var createMetaRepo = func(path string) error {
		admin := jsonMetaRepoAdmin(nil)
		return admin.Create(path)
	}

	BeforeEach(func() {
		testFsRoot = expect.NoError(os.MkdirTemp("", "JsonMetaDataRepo-"))
		DeferCleanup(os.RemoveAll, testFsRoot)
	})

	Describe("#AddLocals", func() {
		It("ignores duplicate paths, given 2 or more of the same exact path", Pending)
		It("ignores duplicate paths, given distinct paths that resolve to the same absolute path", Pending)
		It("accepts absolute paths", Pending)
		It("rejects relative paths", Pending)
	})

	Describe("#AddRemotes", func() {
		It("ignores duplicate URLs, given 2 or more of the same URL", Pending)
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
			Expect(subject.AddLocal("/path/to/one")).To(Succeed())
		})

		It("#AddLocal skips paths that have already been added", Pending, func() {})

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

		It("#AddRemotes skips URLs that have already been added", Pending)

		It("#ListRemote includes each registered remote", func() {
			listTwo := expect.NoError(subject.ListRemote())
			Expect(listTwo.RemoteHrefs()).To(ConsistOf(
				"https://github.com/me/a",
				"https://github.com/me/b",
			))
		})
	})
})
