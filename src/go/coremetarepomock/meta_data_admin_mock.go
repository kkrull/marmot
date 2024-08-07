package coremetarepomock

import (
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Construct a test double for MetaDataAdmin.
func NewMetaDataAdmin() *MetaDataAdmin {
	return &MetaDataAdmin{
		isMetaRepoReturns: make(map[string]bool),
	}
}

// Mock implementation for testing with MetaDataAdmin.
type MetaDataAdmin struct {
	createCalls       []string
	createError       error
	isMetaRepoReturns map[string]bool
}

func (admin *MetaDataAdmin) Create(metaRepoPath string) error {
	admin.createCalls = append(admin.createCalls, metaRepoPath)
	return admin.createError
}

func (admin *MetaDataAdmin) CreateExpected(expectedPath string) {
	ginkgo.GinkgoHelper()
	Expect(admin.createCalls).To(ContainElement(expectedPath))
}

func (admin *MetaDataAdmin) CreateFails(err error) {
	admin.createError = err
}

func (admin *MetaDataAdmin) IsMetaRepo(path string) bool {
	return admin.isMetaRepoReturns[path]
}

func (admin *MetaDataAdmin) IsMetaRepoReturns(path string, value bool) {
	admin.isMetaRepoReturns[path] = value
}
