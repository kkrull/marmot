package svcfs_test

import (
	"io/fs"
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
		It("creates files in the directory, given a valid, writable path", func() {
			subject = jsonMetaRepoAdmin(nil)
			subject.Create(metaRepoPath)

			metaDataDir := filepath.Join(metaRepoPath, ".marmot")
			Expect(os.Stat(metaDataDir)).NotTo(BeNil())

			metaDataFile := filepath.Join(metaDataDir, "meta-repo.json")
			Expect(os.Stat(metaDataFile)).NotTo(BeNil())
		})

		It("returns no error, upon success", func() {
			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create(metaRepoPath)).To(Succeed())
		})

		It("accepts an existing directory that is not a Marmot repo", func() {
			Expect(os.MkdirAll(metaRepoPath, fs.ModePerm)).To(Succeed())
			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create(metaRepoPath)).To(Succeed())
		})

		It("ignores an existing directory already containing Marmot data", func() {
			//What about the files inside of .marmot/?  Should those be re-created or left alone?
			marmotDataDir := filepath.Join(metaRepoPath, ".marmot")
			Expect(os.MkdirAll(marmotDataDir, fs.ModePerm)).To(Succeed())

			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create(metaRepoPath)).To(Succeed())
		})

		It("returns an error, given an invalid path", func() {
			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create("\000x")).To(
				MatchError(MatchRegexp("failed to check for existing meta repo")))
		})

		It("returns an error, given a path in which files can not be created", func() {
			Expect(os.Chmod(testFsRoot, 0o555)).To(Succeed())
			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create(metaRepoPath)).To(
				MatchError(MatchRegexp("failed to make directory")))
		})
	})

	Describe("#IsMetaRepo", func() {
		It("returns false, given a non-existent path", Pending)
		It("returns false, given an existing path that is not a directory", Pending)
		It("returns false, given a directory not containing a Marmot metadata", Pending)
		It("returns true, given a directory containing Marmot metadata", Pending)
	})
})

func jsonMetaRepoAdmin(args *jsonMetaRepoAdminArgs) *svcfs.JsonMetaRepoAdmin {
	if args == nil {
		args = &jsonMetaRepoAdminArgs{}
	}

	return svcfs.NewJsonMetaRepoAdmin(args.Version())
}

type jsonMetaRepoAdminArgs struct {
	version string
}

func (args jsonMetaRepoAdminArgs) Version() string {
	if args.version == "" {
		return "42"
	} else {
		return args.version
	}
}
