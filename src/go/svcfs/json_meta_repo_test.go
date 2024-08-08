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
		admin      *svcfs.JsonMetaRepoAdmin
		testFsRoot string
	)

	BeforeEach(func() {
		testFsRoot = expect.NoError(os.MkdirTemp("", "JsonMetaDataRepo-"))
		DeferCleanup(os.RemoveAll, testFsRoot)
	})

	Describe("#AddRemotes", func() {
		It("adds only distinct URLs, instead of adding duplicates", Pending)
		It("adds only new URLs that don't exist yet", Pending)
	})

	Context("when no repositories have been registered", func() {
		BeforeEach(func() {
			admin = jsonMetaRepoAdmin(nil)
			subject = svcfs.NewJsonMetaRepo(testFsRoot)
			Expect(admin.Create(testFsRoot)).To(Succeed())
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
			admin = jsonMetaRepoAdmin(nil)
			subject = svcfs.NewJsonMetaRepo(testFsRoot)
			Expect(admin.Create(testFsRoot)).To(Succeed())

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
			admin = jsonMetaRepoAdmin(nil)
			subject = svcfs.NewJsonMetaRepo(testFsRoot)
			Expect(admin.Create(testFsRoot)).To(Succeed())

			given := testdata.NewURLs(
				"https://github.com/me/a",
				"https://github.com/me/b",
			)

			Expect(subject.AddRemotes(given)).To(Succeed())
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
