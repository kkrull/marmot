package cmd_test

import (
	"github.com/kkrull/marmot-core/cmd"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("cmd.InitCmd", func() {
	Describe("#Run", func() {
		It("succeeds, given valid conditions", func() {
			fsMock := &MockMetaDataStore{}
			subject := cmd.InitCmd{MetaDataStore: fsMock}
			Expect(subject.Run()).To(Succeed())
		})

		It("ensures there is a place to store meta data", func() {
			dataStoreMock := &MockMetaDataStore{}
			subject := cmd.InitCmd{MetaDataStore: dataStoreMock}

			_ = subject.Run()
			dataStoreMock.EnsureCreatedExpected()
		})

		PIt("initializes meta data, when none exists")

		PIt("does nothing or returns an error when the repository already exists")
	})
})

type MockMetaDataStore struct {
	EnsureCreatedCount int
}

func (fs *MockMetaDataStore) EnsureCreated() {
	fs.EnsureCreatedCount += 1
}

func (fs *MockMetaDataStore) EnsureCreatedExpected() {
	Expect(fs.EnsureCreatedCount).To(BeNumerically(">", 0))
}
