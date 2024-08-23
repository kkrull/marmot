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
	)

	//Paths
	var (
		missingPath = func() string {
			path := filepath.Join(testFsRoot, "missing")
			_, statErr := os.Stat(path)
			Expect(errors.Is(statErr, os.ErrNotExist)).To(BeTrue())
			return path
		}

		validPath = func() string {
			return testFsRoot
		}
	)

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
			Expect(runV(validPath())).To(Succeed())
		})

		//TODO KDK: The paths for the other tests need to exist now
		It("returns an error, given a local path that does not exist", Pending, func() {
			Expect(runV(missingPath())).To(
				MatchError(ContainSubstring("path does not exist")))
		})

		It("returns an error, when adding repositories fails", func() {
			source.AddLocalsFails(errors.New("bang!"))
			Expect(runV(validPath())).To(
				MatchError(ContainSubstring("failed to add local repositories; bang!")))
		})
	})
})
