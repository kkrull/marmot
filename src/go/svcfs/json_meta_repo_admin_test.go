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
		subject    *svcfs.JsonMetaRepoAdmin
		testFsRoot string
	)

	var (
		existingPath    = func() string { return testFsRoot }
		nonExistentPath = func() string { return filepath.Join(testFsRoot, "not-created-yet") }
		someFile        = func() (string, error) {
			path := filepath.Join(testFsRoot, "existing-file")
			if aFile, createErr := os.Create(path); createErr != nil {
				return "", createErr
			} else {
				defer aFile.Close()
				return path, nil
			}
		}
	)

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
			// What about the files inside of .marmot/?  Should those be re-created or left alone?
			marmotDataDir := filepath.Join(testFsRoot, ".marmot")
			Expect(os.MkdirAll(marmotDataDir, fs.ModePerm)).To(Succeed())

			subject = jsonMetaRepoAdmin(nil)
			Expect(subject.Create(testFsRoot)).To(Succeed())
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

		It("returns false, given a directory not containing Marmot metadata", func() {
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
