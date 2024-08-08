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
		subject *userepository.RegisterLocalRepositoriesCommand
		source  *mock.RepositorySource
	)

	var validPaths = func() []string {
		return []string{"/path/to/repo"}
	}

	BeforeEach(func() {
		source = mock.NewRepositorySource()
		factory := use.NewCommandFactory().WithRepositorySource(source)
		subject = expect.NoError(factory.NewRegisterLocalRepositories())
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

		It("returns no error, upon success", func() {
			Expect(subject.Run(validPaths())).To(Succeed())
		})

		It("returns an error, when adding repositories fails", func() {
			source.AddLocalsFails(errors.New("bang!"))
			Expect(subject.Run(validPaths())).To(
				MatchError(ContainSubstring("failed to add local repositories; bang!")))
		})

		It("normalizes paths by replacing redundant/repeated parts with shorter equivalents", func() {
			subject.Run([]string{"/path/to/a/../b"})
			source.AddLocalsExpected("/path/to/b")
		})

		It("resolves relative paths to absolute paths", func() {
			Expect(os.Chdir(os.TempDir())).To(Succeed())
			subject.Run([]string{"some-name-in-tmp"})

			Expect(filepath.Abs("some-name-in-tmp")).To(HaveSuffix(string(os.PathSeparator) + "some-name-in-tmp"))
			source.AddLocalExpectedM(HaveSuffix(string(os.PathSeparator) + "some-name-in-tmp"))
		})

		It("rejects invalid paths", Pending)
		It("rejects paths that do not exist", Pending)
		It("ignores duplicate paths, given distinct paths that resolve to the same absolute path", Pending)
	})
})
