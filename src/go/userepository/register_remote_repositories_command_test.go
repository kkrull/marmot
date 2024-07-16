package userepository_test

import (
	"net/url"

	mock "github.com/kkrull/marmot/corerepositorymock"
	expect "github.com/kkrull/marmot/expect"
	main "github.com/kkrull/marmot/mainfactory"
	"github.com/kkrull/marmot/testdata"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegisterRepositoriesCommand", func() {
	var factory *main.CommandFactory
	var source *mock.RepositorySource

	BeforeEach(func() {
		source = &mock.RepositorySource{}
		factory = &main.CommandFactory{RepositorySource: source}
	})

	Describe("#Run", func() {
		It("succeeds, given valid URLs", func() {
			subject := expect.NoError(factory.RegisterRemoteRepositoriesCommand())
			Expect(subject.Run(validUrls())).To(Succeed())
		})

		It("registers remote repositories at the given URLs", func() {
			subject := expect.NoError(factory.RegisterRemoteRepositoriesCommand())
			Expect(subject.Run(testdata.NewURLs("https://github.com/actions/checkout"))).To(Succeed())
			source.RegisterRemoteExpected("https://github.com/actions/checkout")
		})
	})
})

/* Test data */

func validUrls() []*url.URL {
	return testdata.NewURLs("https://github.com/actions/checkout")
}
