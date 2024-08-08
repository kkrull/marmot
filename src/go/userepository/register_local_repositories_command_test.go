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
		It("passes local paths to the repository source", func() {
			subject.Run("/path/to/a", "/path/to/b")
			source.AddLocalExpected("/path/to/a", "/path/to/b")
		})

		It("returns nil when everything succeeds", func() {
			Expect(subject.Run(validPath())).To(Succeed())
		})

		It("returns informative errors, when they occur", func() {
			source.AddLocalsFails(errors.New("bang!"))
			Expect(subject.Run("/path/to/faulty-repo")).To(
				MatchError(MatchRegexp("failed to add local repositories.*bang!")))
		})

		It("accepts absolute paths", Pending)
		It("resolves relative paths", Pending)
		It("rejects invalid paths", Pending)
		It("rejects paths that do not exist", Pending)
		It("ignores duplicate paths, given distinct paths that resolve to the same absolute path", Pending)
	})
})

func validPath() string {
	return "/path/to/repo"
}
