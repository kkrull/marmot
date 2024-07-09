package cmd_test

import (
	"github.com/kkrull/marmot-core/cmd"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("cmd.InitCmd", func() {
	Describe("#Run", func() {
		It("succeeds, given valid conditions", func() {
			metaDataMock := &MockMetaDataSource{}
			subject := cmd.InitCmd{MetaDataSource: metaDataMock}
			Expect(subject.Run()).To(Succeed())
		})

		It("ensures there is a place to store meta data", func() {
			metaDataMock := &MockMetaDataSource{}
			subject := cmd.InitCmd{MetaDataSource: metaDataMock}

			_ = subject.Run()
			metaDataMock.EnsureCreatedExpected()
		})

		PIt("initializes meta data, when none exists")

		PIt("does nothing or returns an error when the repository already exists")
	})
})

type MockMetaDataSource struct {
	EnsureCreatedCount int
}

func (fs *MockMetaDataSource) EnsureCreated() {
	fs.EnsureCreatedCount += 1
}

func (fs *MockMetaDataSource) EnsureCreatedExpected() {
	Expect(fs.EnsureCreatedCount).To(BeNumerically(">", 0))
}
