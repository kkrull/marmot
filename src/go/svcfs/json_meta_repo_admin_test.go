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

var testFsRoot string

var _ = Describe("JsonMetaRepoAdmin", func() {
	var subject *svcfs.JsonMetaRepoAdmin

	BeforeEach(func() {
		testFsRoot = expect.NoError(os.MkdirTemp("", "JsonMetaDataRepo-"))
		DeferCleanup(os.RemoveAll, testFsRoot)
	})

	Describe("#Create", func() {
		It("creates files in the directory, given a valid, writable path", func() {
			subject = jsonMetaRepoAdmin(nil)
			subject.Create(testFsRoot)

			metaDataDir := filepath.Join(testFsRoot, ".marmot")
			Expect(os.Stat(metaDataDir)).NotTo(BeNil())

			metaDataFile := filepath.Join(metaDataDir, "meta-repo.json")
			Expect(os.Stat(metaDataFile)).NotTo(BeNil())
		})

		It("returns no error, upon success", func() {
			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create(testFsRoot)).To(Succeed())
		})

		It("accepts an existing directory that is not a Marmot repo", func() {
			Expect(os.MkdirAll(testFsRoot, fs.ModePerm)).To(Succeed())
			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create(testFsRoot)).To(Succeed())
		})

		It("ignores an existing directory already containing Marmot data", func() {
			//What about the files inside of .marmot/?  Should those be re-created or left alone?
			marmotDataDir := filepath.Join(testFsRoot, ".marmot")
			Expect(os.MkdirAll(marmotDataDir, fs.ModePerm)).To(Succeed())

			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create(testFsRoot)).To(Succeed())
		})

		It("returns an error, given an invalid path", func() {
			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create("\000x")).To(
				MatchError(MatchRegexp("failed to check for existing meta repo")))
		})

		It("returns an error, given a path in which files can not be created", func() {
			Expect(os.Chmod(testFsRoot, 0o555)).To(Succeed())
			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create(testFsRoot)).To(
				MatchError(MatchRegexp("failed to make directory")))
		})
	})

	Describe("#IsMetaRepo", func() {
		BeforeEach(func() {
			subject = jsonMetaRepoAdmin(nil)
		})

		It("returns false, given a non-existent path", func() {
			Expect(subject.IsMetaRepo(nonExistentPath())).To(Equal(false))
		})

		It("returns false, given an existing path that is not a directory", func() {
			existingFile := expect.NoError(someFile())
			Expect(subject.IsMetaRepo(existingFile)).To(Equal(false))
		})

		It("returns false, given a directory not containing a Marmot metadata", func() {
			Expect(subject.IsMetaRepo(existingPath())).To(Equal(false))
		})

		It("returns true, given a directory containing Marmot metadata", func() {
			marmotDataDir := filepath.Join(testFsRoot, ".marmot")
			Expect(os.MkdirAll(marmotDataDir, fs.ModePerm)).To(Succeed())

			Expect(subject.IsMetaRepo(testFsRoot)).To(Equal(true))
		})

		It("returns an error, when checking that path fails", func() {
			_, err := subject.IsMetaRepo("\000x")
			Expect(err).To(MatchError(ContainSubstring("failed to stat meta repo path")))
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

/* Filesystem */

func existingPath() string { return testFsRoot }
func someFile() (string, error) {
	path := filepath.Join(testFsRoot, "existing-file")
	if aFile, createErr := os.Create(path); createErr != nil {
		return "", createErr
	} else {
		defer aFile.Close()
		return path, nil
	}
}

func nonExistentPath() string { return filepath.Join(testFsRoot, "not-created-yet") }
