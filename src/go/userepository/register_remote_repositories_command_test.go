package userepository_test

import (
	"net/url"

	repomock "github.com/kkrull/marmot/corerepositorymock"
	expect "github.com/kkrull/marmot/expect"
	main "github.com/kkrull/marmot/mainfactory"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegisterRepositoriesCommand", func() {
	var factory *main.CommandFactory
	var source *repomock.RepositorySource

	BeforeEach(func() {
		source = &repomock.RepositorySource{}
		factory = &main.CommandFactory{RepositorySource: source}
	})

	Describe("#Run", func() {
		It("succeeds, given valid URLs", func() {
			subject := expect.NoError(factory.RegisterRemoteRepositoriesCommand())
			Expect(subject.Run(validUrls())).To(Succeed())
		})

		It("registers remote repositories at the given URLs", func() {
			subject := expect.NoError(factory.RegisterRemoteRepositoriesCommand())
			Expect(subject.Run(newURLs("https://github.com/actions/checkout"))).To(Succeed())
			source.RegisterRemoteExpected("https://github.com/actions/checkout")
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
		parsedUrl := expect.NoError(url.Parse(rawUrl))
		parsedUrls[i] = *parsedUrl
	}

	return parsedUrls
}

func validUrls() []url.URL {
	return newURLs("https://github.com/actions/checkout")
}
