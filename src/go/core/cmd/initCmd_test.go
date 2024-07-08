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

		// Test the logic only, using an interface to check for the existence of the directory and data files
		PIt("creates a directory, when none exists")
		PIt("initializes meta data, when none exists")
	})
})
