package userepository_test

import (
	"errors"

	mock "github.com/kkrull/marmot/corerepositorymock"
	expect "github.com/kkrull/marmot/testsupportexpect"
	"github.com/kkrull/marmot/use"
	"github.com/kkrull/marmot/userepository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegisterLocalRepositoriesCommand", func() {
	var (
		factory use.CommandFactory
		source  *mock.RepositorySource
		subject *userepository.RegisterLocalRepositoriesCommand
	)

	BeforeEach(func() {
		source = mock.NewRepositorySource()
		factory = use.NewCommandFactory().WithRepositorySource(source)
		subject = expect.NoError(factory.NewRegisterLocalRepositories())
	})

	Describe("#Run", func() {
		It("adds local repositories to the source, for the given paths", func() {
			subject.Run([]string{"/path/to/a", "/path/to/b"})
			source.AddLocalsExpected("/path/to/a", "/path/to/b")
		})

		It("returns no error, upon success", func() {
			Expect(subject.Run(validPath())).To(Succeed())
		})

		It("returns an error, when adding repositories fails", func() {
			source.AddLocalsFails(errors.New("bang!"))
			Expect(subject.Run([]string{"/path/to/faulty-repo"})).To(
				MatchError(ContainSubstring("failed to add local repositories; bang!")))
		})

		It("accepts absolute paths", Pending)
		It("resolves relative paths", Pending)
		It("rejects invalid paths", Pending)
		It("rejects paths that do not exist", Pending)
		It("ignores duplicate paths, given distinct paths that resolve to the same absolute path", Pending)
	})
})

func validPath() []string {
	return []string{"/path/to/repo"}
}
