package cmd_test

import (
	"github.com/kkrull/marmot-core/cmd"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("cmd.InitCmd", func() {
	Describe("#Run", func() {
		It("succeeds, given valid conditions", func() {
			fsMock := &MockMarmotFileSystem{}
			subject := cmd.InitCmd{FileSystem: fsMock}
			Expect(subject.Run()).To(Succeed())
		})

		It("creates a directory, when none exists", func() {
			fsMock := &MockMarmotFileSystem{}
			subject := cmd.InitCmd{FileSystem: fsMock}

			_ = subject.Run()
			fsMock.EnsureExistsExpected("/path/to/meta")
		})

		PIt("initializes meta data, when none exists")

		PIt("does nothing or returns an error when the repository already exists")
	})
})

type MockMarmotFileSystem struct {
	EnsureExistsReceived []string
}

func (fs *MockMarmotFileSystem) EnsureExists(path string) {
	fs.EnsureExistsReceived = append(fs.EnsureExistsReceived, path)
}

func (fs *MockMarmotFileSystem) EnsureExistsExpected(expectedPath string) {
	GinkgoHelper()
	Expect(fs.EnsureExistsReceived).To(ContainElement(expectedPath))
}
