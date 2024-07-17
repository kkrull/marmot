package coremetarepomock

import (
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type MetaDataAdmin struct {
	InitCalls []string
	InitCount int
	InitError error
}

func (admin *MetaDataAdmin) Init(metaRepoPath string) error {
	admin.InitCalls = append(admin.InitCalls, metaRepoPath)
	admin.InitCount += 1
	return admin.InitError
}

func (admin *MetaDataAdmin) InitExpected(expectedPath string) {
	ginkgo.GinkgoHelper()
	Expect(admin.InitCalls).To(ContainElement(expectedPath))
}
