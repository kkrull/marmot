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
		dirFixture *testsupportfs.DirFixture
		source     *mock.RepositorySource
		subject    *userepository.RegisterLocalRepositoriesCommand
	)

	var validPaths = func() []string {
		return []string{"/path/to/repo"}
	}

	BeforeEach(func() {
		dirFixture = testsupportfs.NewDirFixture("RegisterLocalRepositoriesCommand")
		Expect(dirFixture.Setup()).To(Succeed())
		DeferCleanup(dirFixture.Teardown)

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
			subject.Run([]string{"some-name-in-cwd"})
			Expect(source.AddLocalsReceived()[0]).To(
				HaveSuffix(string(os.PathSeparator) + "some-name-in-cwd"))
		})
	})
})
