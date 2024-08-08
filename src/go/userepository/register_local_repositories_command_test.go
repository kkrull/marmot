package userepository_test

import (
	"errors"
	"os"

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

		It("ignores distinct paths that resolve to an already-known absolute path", func() {
			subject.Run([]string{"somewhere", "./somewhere"})
			Expect(source.AddLocalsReceived()).To(HaveLen(1))
			Expect(source.AddLocalsReceived()[0]).To(HaveSuffix("somewhere"))
		})

		It("normalizes paths by replacing redundant/repeated parts with shorter equivalents", func() {
			subject.Run([]string{"/path/to/a/../b"})
			source.AddLocalsExpected("/path/to/b")
		})

		It("resolves relative paths to absolute paths", func() {
			subject.Run([]string{"some-name-in-cwd"})
			source.AddLocalExpectedM(HaveSuffix(string(os.PathSeparator) + "some-name-in-cwd"))
		})

		It("rejects invalid paths", Pending)
		It("rejects paths that do not exist", Pending)
	})
})
