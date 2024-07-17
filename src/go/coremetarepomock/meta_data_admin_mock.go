package coremetarepomock

import (
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type MetaDataAdmin struct {
	CreateCalls []string
	CreateError error
}

func (admin *MetaDataAdmin) Create(metaRepoPath string) error {
	admin.CreateCalls = append(admin.CreateCalls, metaRepoPath)
	return admin.CreateError
}

func (admin *MetaDataAdmin) CreateExpected(expectedPath string) {
	ginkgo.GinkgoHelper()
	Expect(admin.CreateCalls).To(ContainElement(expectedPath))
}
