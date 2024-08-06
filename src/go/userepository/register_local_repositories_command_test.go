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
			source.AddLocalFails("/path/to/faulty-repo", errors.New("bang!"))
			Expect(subject.Run("/path/to/faulty-repo")).To(MatchError("failed to add local repository /path/to/faulty-repo: bang!"))
		})

		It("rejects invalid paths", Pending, func() {})
		It("rejects paths that do not exist", Pending, func() {})
		It("resolves relative paths", Pending, func() {})
	})
})

func validPath() string {
	return "/path/to/repo"
}
