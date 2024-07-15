package coremetarepomock

import (
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type MetaDataAdmin struct {
	InitCount int
	InitError error
}

func (fs *MetaDataAdmin) Init() error {
	fs.InitCount += 1
	return fs.InitError
}

func (fs *MetaDataAdmin) InitExpected() {
	ginkgo.GinkgoHelper()
	Expect(fs.InitCount).To(Equal(1))
}
