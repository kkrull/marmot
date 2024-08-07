package usemetarepo_test

import (
	"errors"
	"os"
	"path/filepath"

	mock "github.com/kkrull/marmot/coremetarepomock"
	expect "github.com/kkrull/marmot/testsupportexpect"
	"github.com/kkrull/marmot/use"
	"github.com/kkrull/marmot/usemetarepo"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var testDir string

func existingPath() string    { return testDir }
func nonExistentPath() string { return filepath.Join(testDir, "not-created-yet") }
func validPath() string       { return filepath.Join(testDir, "meta-default") }

var _ = Describe("InitCommand", func() {
	var (
		subject       *usemetarepo.InitCommand
		metaDataAdmin *mock.MetaDataAdmin
	)

	BeforeEach(func() {
		testDir = expect.NoError(os.MkdirTemp("", "InitCommand-"))
		DeferCleanup(os.RemoveAll, testDir)

		metaDataAdmin = mock.NewMetaDataAdmin()
		factory := use.NewCommandFactory().WithMetaDataAdmin(metaDataAdmin)
		subject = expect.NoError(factory.NewInitMetaRepo())
	})

	Describe("#Run", func() {
		It("creates a meta repo in that path", func() {
			givenPath := validPath()
			Expect(subject.Run(givenPath)).To(Succeed())
			metaDataAdmin.CreateExpected(givenPath)
		})

		It("returns a nil error, when that succeeds", func() {
			Expect(subject.Run(validPath())).To(Succeed())
		})

		It("accepts paths that do and do not exist, provided a meta repo is not there", func() {
			Expect(subject.Run(existingPath())).To(Succeed())
			Expect(subject.Run(nonExistentPath())).To(Succeed())
		})

		It("returns an error when the path is already a meta repo", func() {
			existingMetaRepo := filepath.Join(testDir, "meta-already")
			metaDataAdmin.ExistsReturns(existingMetaRepo, true)

			Expect(subject.Run(existingMetaRepo)).To(
				MatchError(MatchRegexp("meta-already is already a meta repo$")))
		})

		It("returns an error when creating a meta repo fails", func() {
			metaDataAdmin.CreateFails(errors.New("bang!"))
			Expect(subject.Run(validPath())).To(
				MatchError(MatchRegexp("^failed to initialize meta repo.*bang!$")))
		})
	})
})
