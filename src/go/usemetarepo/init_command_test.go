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
		subject        *metarepo.InitCommand
		cmdFactory     *main.CommandFactory
		metaDataSource *MockMetaDataSource
	)

	BeforeEach(func() {
		metaDataSource = &MockMetaDataSource{}
		cmdFactory = &main.CommandFactory{MetaDataSource: metaDataSource}
	})

	Describe("#Run", func() {
		It("initializes the given meta data source", func() {
			subject, _ = cmdFactory.InitCommand()
			_ = subject.Run()
			metaDataSource.InitExpected()
		})

		It("returns nil, when everything succeeds", func() {
			subject, _ = cmdFactory.InitCommand()
			Expect(subject.Run()).To(BeNil())
		})

		It("returns an error when failing to initialize the meta data source", func() {
			metaDataSource.InitError = errors.New("bang!")

			subject, _ = cmdFactory.InitCommand()
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
