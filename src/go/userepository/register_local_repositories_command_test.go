package userepository_test

import (
	"errors"
	"os"

	mock "github.com/kkrull/marmot/corerepositorymock"
	expect "github.com/kkrull/marmot/testsupportexpect"
	"github.com/kkrull/marmot/testsupportfs"
	"github.com/kkrull/marmot/use"
	"github.com/kkrull/marmot/userepository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegisterLocalRepositoriesCommand", func() {
	var (
		subject *userepository.RegisterLocalRepositoriesCommand
		runV    = func(paths ...string) error {
			return subject.Run(paths)
		}
	)

	var source *mock.RepositorySource

	var (
		dirFixture  *testsupportfs.DirFixture
		pathBuilder *testsupportfs.PathBuilder
	)

	var validPaths = func() []string {
		return []string{"/path/to/repo"}
	}

	BeforeEach(func() {
		dirFixture = testsupportfs.NewDirFixture("RegisterLocalRepositoriesCommand")
		Expect(dirFixture.Setup()).To(Succeed())
		DeferCleanup(dirFixture.Teardown)
		pathBuilder = expect.NoError(dirFixture.PathBuilder())

		source = mock.NewRepositorySource()
		factory := use.NewCommandFactory().WithRepositorySource(source)
		subject = expect.NoError(factory.NewRegisterLocalRepositories())
	})

	Describe("#Run", func() {
		It("accepts absolute paths", func() {
			Expect(runV(pathBuilder.AsAbsolute())).To(Succeed())
		})

		It("accepts relative paths", func() {
			Expect(subject.Run([]string{"../git/relative"})).To(Succeed())
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
			subject.Run([]string{"some-name-in-cwd"})
			Expect(source.AddLocalsReceived()[0]).To(
				HaveSuffix(string(os.PathSeparator) + "some-name-in-cwd"))
		})
	})
})
