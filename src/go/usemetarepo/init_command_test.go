package usemetarepo_test

import (
	"errors"
	"os"

	mock "github.com/kkrull/marmot/coremetarepomock"
	expect "github.com/kkrull/marmot/testsupportexpect"
	"github.com/kkrull/marmot/use"
	"github.com/kkrull/marmot/usemetarepo"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("InitCommand", func() {
	var (
		subject       *usemetarepo.InitCommand
		metaDataAdmin *mock.MetaDataAdmin
		testDir       string
	)

	BeforeEach(func() {
		testDir = expect.NoError(os.MkdirTemp("", "InitCommand-"))
		DeferCleanup(os.RemoveAll, testDir)

		metaDataAdmin = mock.NewMetaDataAdmin()
		factory := use.NewCommandFactory().WithMetaDataAdmin(metaDataAdmin)
		subject = expect.NoError(factory.NewInitMetaRepo())
	})

	Describe("#Run (new interface)", func() {
		It("creates a meta repo in that path", Pending, func() {})
		It("returns a nil error, when that succeeds", Pending, func() {})
		It("allows a meta repo to be created in an existing directory", Pending, func() {})
		It("returns an error when the path is already a meta repo", Pending, func() {})
		It("returns an error when creating a meta repo fails", Pending, func() {})
	})

	Describe("#Run (old interface)", func() {
		It("returns an error when t˝he path already contains a meta repo", Pending, func() {
			metaDataAdmin.ExistsReturns(testDir, true)

			Expect(subject.Run(testDir)).To(
				MatchError(MatchRegexp("^InitCommand.*: already a meta repo$")))
		})

		It("initializes the given meta data source", func() {
			_ = subject.Run("/tmp")
			metaDataAdmin.CreateExpected("/tmp")
		})

		It("returns nil, when everything succeeds", func() {
			Expect(subject.Run("/tmp")).To(BeNil())
		})

		It("returns an error when failing to initialize the meta data source", func() {
			metaDataAdmin.CreateFails(errors.New("bang!"))
			Expect(subject.Run("/tmp")).To(MatchError("bang!"))
		})
	})
})
