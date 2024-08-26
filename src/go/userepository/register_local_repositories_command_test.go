package userepository_test

import (
	"errors"
	"os"

	mock "github.com/kkrull/marmot/corerepositorymock"
	"github.com/kkrull/marmot/testsupportdata"
	expect "github.com/kkrull/marmot/testsupportexpect"
	"github.com/kkrull/marmot/use"
	"github.com/kkrull/marmot/userepository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegisterLocalRepositoriesCommand", func() {
	//Subject, dependencies, and convenience functions
	var (
		source  *mock.RepositorySource
		subject *userepository.RegisterLocalRepositoriesCommand
		runV    = func(localPaths ...string) error {
			return subject.Run(localPaths)
		}
	)

	//Test fixture
	var (
		originalCwd string
		testFsRoot  string
		pathFixture *testsupportdata.PathBuilder
	)

	BeforeEach(func() {
		source = mock.NewRepositorySource()
		factory := use.NewCommandFactory().WithRepositorySource(source)
		subject = expect.NoError(factory.NewRegisterLocalRepositories())

		pathFixture = testsupportdata.NewPathBuilder("RegisterLocalRepositoriesCommand")
		Expect(pathFixture.Setup()).To(Succeed())
		DeferCleanup(pathFixture.Teardown())
	})

	Describe("#Run", func() {
		//TODO KDK: Start expending PathBuilder to take in properties like absolute/relative, existing/not, git repo/not
		It("accepts absolute paths", func() {
			Expect(runV("/home/me/absolute")).To(Succeed())
		})

		It("accepts relative paths", func() {
			Expect(runV("../git/relative")).To(Succeed())
		})

		It("adds local repositories to the source, for the given paths", func() {
			runV("/path/to/a", "/path/to/b")
			source.AddLocalsExpected("/path/to/a", "/path/to/b")
		})

		It("normalizes paths by replacing redundant/repeated parts with shorter equivalents", func() {
			runV("/path/to/a/../b")
			source.AddLocalsExpected("/path/to/b")
		})

		It("resolves relative paths to absolute paths", func() {
			runV("some-name-in-cwd")
			Expect(source.AddLocalsReceived()[0]).To(
				HaveSuffix(string(os.PathSeparator) + "some-name-in-cwd"))
		})

		It("returns no error, upon success", func() {
			path := pathFixture.Build()
			Expect(runV(path)).To(Succeed())
		})

		//TODO KDK: The paths for the other tests need to exist now
		It("returns an error, given a local path that does not exist", Pending, func() {
			path := pathFixture.Missing().Build()
			Expect(runV(path)).To(
				MatchError(ContainSubstring("path does not exist")))
		})

		It("returns an error, when adding repositories fails", func() {
			path := pathFixture.Build()
			source.AddLocalsFails(errors.New("bang!"))
			Expect(runV(path)).To(
				MatchError(ContainSubstring("failed to add local repositories; bang!")))
		})
	})
})
