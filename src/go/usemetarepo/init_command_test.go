package usemetarepo_test

import (
	"errors"
	"path/filepath"

	mock "github.com/kkrull/marmot/coremetarepomock"
	expect "github.com/kkrull/marmot/testsupportexpect"
	"github.com/kkrull/marmot/testsupportfs"
	"github.com/kkrull/marmot/use"
	"github.com/kkrull/marmot/usemetarepo"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var testFsRoot string

func existingPath() string    { return testFsRoot }
func nonExistentPath() string { return filepath.Join(testFsRoot, "not-created-yet") }
func validPath() string       { return filepath.Join(testFsRoot, "meta-default") }

var _ = Describe("InitCommand", func() {
	var (
		subject       *usemetarepo.InitCommand
		dirFixture    *testsupportfs.DirFixture
		metaDataAdmin *mock.MetaDataAdmin
	)

	BeforeEach(func() {
		dirFixture = testsupportfs.NewDirFixture("JsonMetaDataRepo-")
		Expect(dirFixture.Setup()).To(Succeed())
		DeferCleanup(dirFixture.Teardown)
		testFsRoot = expect.NoError(dirFixture.BasePath())

		metaDataAdmin = mock.NewMetaDataAdmin()
		factory := use.NewCommandFactory().WithMetaDataAdmin(metaDataAdmin)
		subject = expect.NoError(factory.NewInitMetaRepo())
	})

	Describe("#Run", func() {
		It("creates a meta repo in that path", func() {
			givenPath := validPath()
			subject.Run(givenPath)
			metaDataAdmin.CreateExpected(givenPath)
		})

		It("returns no error, upon success", func() {
			Expect(subject.Run(validPath())).To(Succeed())
		})

		It("accepts paths that do and do not exist, provided a meta repo is not there", func() {
			Expect(subject.Run(existingPath())).To(Succeed())
			Expect(subject.Run(nonExistentPath())).To(Succeed())
		})

		It("returns an error when unable to check the path", func() {
			path := filepath.Join(testFsRoot, "stealth")
			metaDataAdmin.IsMetaRepoReturns(path, false, errors.New("bang!"))
			Expect(subject.Run(path)).To(
				MatchError(ContainSubstring("stealth: unable to check path; bang!")))
		})

		It("returns an error when the path is already a meta repo", func() {
			existingMetaRepo := filepath.Join(testFsRoot, "meta-already")
			metaDataAdmin.IsMetaRepoReturns(existingMetaRepo, true, nil)
			Expect(subject.Run(existingMetaRepo)).To(
				MatchError(MatchRegexp("meta-already: already a meta repo$")))
		})

		It("returns an error when creating a meta repo fails", func() {
			metaDataAdmin.CreateFails(errors.New("bang!"))
			Expect(subject.Run(validPath())).To(
				MatchError(MatchRegexp("^failed to initialize meta repo.*bang!$")))
		})
	})
})
