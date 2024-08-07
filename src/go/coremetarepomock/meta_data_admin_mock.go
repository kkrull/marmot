package coremetarepomock

import (
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Construct a test double for MetaDataAdmin.
func NewMetaDataAdmin() *MetaDataAdmin {
	return &MetaDataAdmin{}
}

// Mock implementation for testing with MetaDataAdmin.
type MetaDataAdmin struct {
	createCalls []string
	createError error
}

func (admin *MetaDataAdmin) Create(metaRepoPath string) error {
	admin.createCalls = append(admin.createCalls, metaRepoPath)
	return admin.createError
}

// Assert that a meta repo was created at the specified path.
func (admin *MetaDataAdmin) CreateExpected(expectedPath string) {
	ginkgo.GinkgoHelper()
	Expect(admin.createCalls).To(ContainElement(expectedPath))
}

// Stub #Create to fail with the given error.
func (admin *MetaDataAdmin) CreateFails(err error) {
	admin.createError = err
}

func (admin *MetaDataAdmin) ExistsReturns(path string, value bool) {
	//TODO KDK: Implement
}
