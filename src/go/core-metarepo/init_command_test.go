package core_metarepo_test

import (
	"errors"

	metarepo "github.com/kkrull/marmot/core-metarepo"
	factory "github.com/kkrull/marmot/main-factory"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("InitCommand", func() {
	var subject *metarepo.InitCommand
	var cmdFactory *factory.CommandFactory
	var metaDataSource *MockMetaDataSource

	BeforeEach(func ()  {
			metaDataSource = &MockMetaDataSource{}
			cmdFactory = &factory.CommandFactory{MetaDataSource: metaDataSource}
	})

	Describe("#Run", func() {
		It("initializes the given meta data source", func() {
			subject, _ = cmdFactory.NewInitCommand()
			_ = subject.Run()
			metaDataSource.InitExpected()
		})

		It("returns nil, when everything succeeds", func() {
			subject, _ = cmdFactory.NewInitCommand()
			Expect(subject.Run()).To(BeNil())
		})

		It("returns an error when failing to initialize the meta data source", func() {
			metaDataSource.InitError = errors.New("bang!")

			subject, _ = cmdFactory.NewInitCommand()
			Expect(subject.Run()).To(MatchError("bang!"))
		})
	})
})

type MockMetaDataSource struct {
	InitCount int
	InitError error
}

func (fs *MockMetaDataSource) Init() error {
	fs.InitCount += 1
	return fs.InitError
}

func (fs *MockMetaDataSource) InitExpected() {
	Expect(fs.InitCount).To(Equal(1))
}
