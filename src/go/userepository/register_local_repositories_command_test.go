package userepository_test

import (
	"errors"
	"os"
	"path/filepath"

	mock "github.com/kkrull/marmot/corerepositorymock"
	expect "github.com/kkrull/marmot/testsupportexpect"
	"github.com/kkrull/marmot/use"
	"github.com/kkrull/marmot/userepository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegisterLocalRepositoriesCommand", func() {
	var (
		subject     *userepository.RegisterLocalRepositoriesCommand
		originalCwd string
		source      *mock.RepositorySource
		testFsRoot  string
	)

	var validPaths = func() []string {
		return []string{"/path/to/repo"}
	}

	BeforeEach(func() {
		source = mock.NewRepositorySource()
		factory := use.NewCommandFactory().WithRepositorySource(source)
		subject = expect.NoError(factory.NewRegisterLocalRepositories())

		testFsRoot = expect.NoError(os.MkdirTemp("", "RegisterLocalRepositoriesCommand-"))
		DeferCleanup(os.RemoveAll, testFsRoot)

		originalCwd = expect.NoError(os.Getwd())
		Expect(os.Chdir(testFsRoot)).To(Succeed())
		DeferCleanup(func() error { return os.Chdir(originalCwd) })
	})

	Describe("#Run", func() {
		It("accepts absolute and relative paths", func() {
			Expect(subject.Run([]string{
				"/home/me/absolute",
				"../git/relative",
			})).To(Succeed())
		})

		It("adds local repositories to the source, for the given paths", func() {
			subject.Run([]string{"/path/to/a", "/path/to/b"})
			source.AddLocalsExpected("/path/to/a", "/path/to/b")
		})

		It("normalizes paths by replacing redundant/repeated parts with shorter equivalents", func() {
			subject.Run([]string{"/path/to/a/../b"})
			source.AddLocalsExpected("/path/to/b")
		})

		It("resolves relative paths to absolute paths", func() {
			subject.Run([]string{"some-name-in-cwd"})
			Expect(source.AddLocalsReceived()[0]).To(
				HaveSuffix(string(os.PathSeparator) + "some-name-in-cwd"))
		})

		It("returns no error, upon success", func() {
			Expect(subject.Run(validPaths())).To(Succeed())
		})

		It("returns an error, given a local path that does not exist", func() {
			missingPath := filepath.Join(testFsRoot, "missing")
			_, statErr := os.Stat(missingPath)
			Expect(errors.Is(statErr, os.ErrNotExist)).To(BeTrue())

			Expect(subject.Run([]string{missingPath})).To(
				MatchError(ContainSubstring("path does not exist")))
		})

		It("returns an error, when adding repositories fails", func() {
			source.AddLocalsFails(errors.New("bang!"))
			Expect(subject.Run(validPaths())).To(
				MatchError(ContainSubstring("failed to add local repositories; bang!")))
		})
	})
})
