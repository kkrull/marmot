package cmd_test

import (
	"github.com/kkrull/marmot-core/cmd"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("cmd.InitCmd", func() {
	Describe("#Run", func() {
		It("exists", func() {
			subject := cmd.InitCmd{}
			Expect(subject.Run()).To(Succeed())
		})
	})
})
