package svcfs_test

import (
	"fmt"
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
		It("is cool with an existing path in which marmot has not been initialized", func() {
			Expect(os.MkdirAll(metaRepoPath, fs.ModePerm)).To(Succeed())

			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create(metaRepoPath)).To(Succeed())
		})

		It("returns an error, given a path containing marmot data", Pending, func() {
			marmotDataDir := filepath.Join(metaRepoPath, ".marmot")
			Expect(os.MkdirAll(marmotDataDir, fs.ModePerm)).To(Succeed())

			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create(metaRepoPath)).To(
				MatchError(fmt.Sprintf("path already exists: %s", marmotDataDir)))
		})

		It("returns an error when unable to check if the path exists", func() {
			subject = jsonMetaRepoAdmin(nil)
			invalidPathErr := subject.Create("\000x")
			Expect(invalidPathErr).NotTo(BeNil())
		})

		It("returns an error when creating files fails", func() {
			Expect(os.Chmod(testFsRoot, 0o555)).To(Succeed())

			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create(metaRepoPath)).To(
				MatchError(ContainSubstring(fmt.Sprintf("failed to make directory %s", metaRepoPath))))
		})

		It("creates files in the meta repository and returns nil, otherwise", func() {
			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create(metaRepoPath)).To(Succeed())

			metaDataDir := filepath.Join(metaRepoPath, ".marmot")
			Expect(os.Stat(metaDataDir)).NotTo(BeNil())

			metaDataFile := filepath.Join(metaDataDir, "meta-repo.json")
			Expect(os.Stat(metaDataFile)).NotTo(BeNil())
		})
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
