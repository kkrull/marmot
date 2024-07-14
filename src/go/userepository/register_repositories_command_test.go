package userepository_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegisterRepositoriesCommand", func() {
	Describe("#Run", func() {
		It("registers the given paths as local repositories, given paths to Git repositories", func() {
			Expect("pending").To(Equal("passing"))
		})

		It("registers no remote URLs for a repository, given a Git repository with no remotes", Pending, func() {
		})

		It("registers the URL of each remote for a Git repository, given one with remotes", Pending, func() {
		})
	})
})
