package cmd_test

import (
	"errors"

	"github.com/kkrull/marmot-core/cmd"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("cmd.InitCmd", func() {
	Describe("#Run", func() {
		It("initializes the given meta data source", func() {
			metaDataSourceMock := &MockMetaDataSource{}
			subject := cmd.InitCmd{MetaDataSource: metaDataSourceMock}

			_ = subject.Run()
			metaDataSourceMock.InitExpected()
		})

		It("returns nil, when everything succeeds", func() {
			metaDataSourceMock := &MockMetaDataSource{}
			subject := cmd.InitCmd{MetaDataSource: metaDataSourceMock}
			Expect(subject.Run()).To(BeNil())
		})

		It("returns an error when failing to initialize the meta data source", func() {
			metaDataSourceMock := &MockMetaDataSource{InitError: errors.New("bang!")}
			subject := cmd.InitCmd{MetaDataSource: metaDataSourceMock}
			Expect(subject.Run()).To(MatchError("bang!"))
		})
	})
})

var _ = Describe("fs.JsonMetaDataSource", func() {
	Describe("#Init", func() {
		PIt("returns an error if the directory already exists")
		PIt("returns an error when creating files fails")
		PIt("returns nil, after creating files in a directory that does not exist")
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
