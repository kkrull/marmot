package svcfs_test

import (
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
		subject      *svcfs.JsonMetaRepo
		admin        *svcfs.JsonMetaRepoAdmin
		metaRepoPath string
		testFsRoot   string
	)

	BeforeEach(func() {
		testFsRoot = expect.NoError(os.MkdirTemp("", "JsonMetaDataRepo-"))
		metaRepoPath = filepath.Join(testFsRoot, "meta")
		DeferCleanup(os.RemoveAll, testFsRoot)
	})

	It("#AddLocal exists", func() {
		admin = jsonMetaRepoAdmin(nil)
		subject = svcfs.NewJsonMetaRepo(metaRepoPath)
		Expect(admin.Create(metaRepoPath)).To(Succeed())

		Expect(subject.AddLocal("/path/to/repo")).To(Succeed())
	})

	It("#ListLocal exists", Pending, func() {})

	Context("when no repositories have been registered", func() {
		BeforeEach(func() {
			admin = jsonMetaRepoAdmin(nil)
			subject = svcfs.NewJsonMetaRepo(metaRepoPath)
			Expect(admin.Create(metaRepoPath)).To(Succeed())
		})

		It("#ListRemote returns empty", func() {
			repositories := expect.NoError(subject.ListRemote())
			Expect(repositories.Count()).To(Equal(0))
		})
	})

	Context("when remote repositories have been registered", func() {
		BeforeEach(func() {
			admin = jsonMetaRepoAdmin(nil)
			subject = svcfs.NewJsonMetaRepo(metaRepoPath)
			Expect(admin.Create(metaRepoPath)).To(Succeed())

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
