package usemetarepo_test

import (
	"errors"

	main "github.com/kkrull/marmot/mainfactory"
	metarepo "github.com/kkrull/marmot/usemetarepo"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("InitCommand", func() {
	var (
		subject       *metarepo.InitCommand
		factory       *main.CommandFactory
		metaDataAdmin *MockMetaDataAdmin
	)

	BeforeEach(func() {
		metaDataAdmin = &MockMetaDataAdmin{}
		factory = &main.CommandFactory{MetaDataAdmin: metaDataAdmin}
	})

	Describe("#Run", func() {
		It("initializes the given meta data source", func() {
			subject, _ = factory.InitCommand()
			_ = subject.Run()
			metaDataAdmin.InitExpected()
		})

		It("returns nil, when everything succeeds", func() {
			subject, _ = factory.InitCommand()
			Expect(subject.Run()).To(BeNil())
		})

		It("returns an error when failing to initialize the meta data source", func() {
			metaDataAdmin.InitError = errors.New("bang!")

			subject, _ = factory.InitCommand()
			Expect(subject.Run()).To(MatchError("bang!"))
		})
	})
})

type MockMetaDataAdmin struct {
	InitCount int
	InitError error
}

func (fs *MockMetaDataAdmin) Init() error {
	fs.InitCount += 1
	return fs.InitError
}

func (fs *MockMetaDataAdmin) InitExpected() {
	Expect(fs.InitCount).To(Equal(1))
}
