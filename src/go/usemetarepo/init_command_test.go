package usemetarepo_test

import (
	"errors"

	mock "github.com/kkrull/marmot/coremetarepomock"
	main "github.com/kkrull/marmot/mainfactory"
	use "github.com/kkrull/marmot/usemetarepo"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("InitCommand", func() {
	var (
		subject       *use.InitCommand
		factory       *main.CommandFactory
		metaDataAdmin *mock.MetaDataAdmin
	)

	BeforeEach(func() {
		metaDataAdmin = &mock.MetaDataAdmin{}
		factory = &main.CommandFactory{MetaDataAdmin: metaDataAdmin}
	})

	Describe("#Run", func() {
		It("initializes the given meta data source", func() {
			subject, _ = factory.InitCommand()
			_ = subject.RunP("/tmp")
			metaDataAdmin.InitExpected()
		})

		It("returns nil, when everything succeeds", func() {
			subject, _ = factory.InitCommand()
			Expect(subject.RunP("/tmp")).To(BeNil())
		})

		It("returns an error when failing to initialize the meta data source", func() {
			metaDataAdmin.InitError = errors.New("bang!")

			subject, _ = factory.InitCommand()
			Expect(subject.RunP("/tmp")).To(MatchError("bang!"))
		})
	})
})
