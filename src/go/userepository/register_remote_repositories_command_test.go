package userepository_test

import (
	"net/url"

	main "github.com/kkrull/marmot/mainfactory"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegisterRepositoriesCommand", func() {
	var factory *main.CommandFactory

	BeforeEach(func() {
		factory = &main.CommandFactory{}
	})

	Describe("#Run", func() {
		It("succeeds, given valid URLs", func() {
			subject := ExpectNoError(factory.RegisterRemoteRepositoriesCommand())
			Expect(subject.Run(validUrls())).To(Succeed())
		})

		It("registers remote repositories at the given URLs", func() {
			subject := ExpectNoError(factory.RegisterRemoteRepositoriesCommand())
			Expect(subject.Run(newURLs("https://github.com/actions/checkout"))).To(Succeed())
		})

		It("registers no remote URLs for a repository, given a Git repository with no remotes", Pending, func() {
		})

		It("registers the URL of each remote for a Git repository, given one with remotes", Pending, func() {
		})
	})
})

/* Test data */

func newURLs(rawUrls ...string) []url.URL {
	GinkgoHelper()
	parsedUrls := make([]url.URL, len(rawUrls))
	for i, rawUrl := range rawUrls {
		parsedUrl := ExpectNoError(url.Parse(rawUrl))
		parsedUrls[i] = *parsedUrl
	}

	return parsedUrls
}

func validUrls() []url.URL {
	return newURLs("https://github.com/actions/checkout")
}

/* Test helpers */

// Expect that a value (or not) was returned without an error, then carry on with(out) it.
func ExpectNoError[V any](maybeValue V, unexpectedErr error) V {
	GinkgoHelper()
	Expect(unexpectedErr).To(BeNil())
	return maybeValue
}
