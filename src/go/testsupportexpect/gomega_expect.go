package testsupportexpect

import (
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Expect that a value (or not) was returned without an error, then carry on with(out) it.
func NoError[V any](maybeValue V, unexpectedErr error) V {
	ginkgo.GinkgoHelper()
	Expect(unexpectedErr).To(BeNil())
	return maybeValue
}
