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

func (admin *MetaDataAdmin) InitP(metaRepoPath string) error {
	admin.InitCalls = append(admin.InitCalls, metaRepoPath)
	admin.InitCount += 1
	return admin.InitError
}

func (admin *MetaDataAdmin) InitExpected() {
	ginkgo.GinkgoHelper()
	Expect(admin.InitCount).To(Equal(1))
}

func (admin *MetaDataAdmin) InitPExpected(expectedPath string) {
	ginkgo.GinkgoHelper()
	Expect(admin.InitCalls).To(ContainElement(expectedPath))
}
