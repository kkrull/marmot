package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("core", func() {
	Describe("main", func() {
		It("exists", func() {
			answer := 42
			Expect(answer).To(Equal(42))
		})
	})
})
