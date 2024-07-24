package usemetarepo_test

import (
	"errors"

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
	)

	BeforeEach(func() {
		metaDataAdmin = &mock.MetaDataAdmin{}
		factory := use.NewAppFactory().WithMetaDataAdmin(metaDataAdmin)
		subject = expect.NoError(factory.InitCommand())
	})

	Describe("#Run", func() {
		It("initializes the given meta data source", func() {
			_ = subject.Run("/tmp")
			metaDataAdmin.CreateExpected("/tmp")
		})

		It("returns nil, when everything succeeds", func() {
			Expect(subject.Run("/tmp")).To(BeNil())
		})

		It("returns an error when failing to initialize the meta data source", func() {
			metaDataAdmin.CreateError = errors.New("bang!")
			Expect(subject.Run("/tmp")).To(MatchError("bang!"))
		})
	})
})
