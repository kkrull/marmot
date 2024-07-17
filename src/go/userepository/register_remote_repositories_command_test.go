package userepository_test

import (
	"net/url"

	mock "github.com/kkrull/marmot/corerepositorymock"
	main "github.com/kkrull/marmot/mainfactory"
	testdata "github.com/kkrull/marmot/testsupportdata"
	expect "github.com/kkrull/marmot/testsupportexpect"
	"github.com/kkrull/marmot/userepository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegisterRepositoriesCommand", func() {
	var (
		factory *main.CommandFactory
		source  *mock.RepositorySource
		subject *userepository.RegisterRemoteRepositoriesCommand
	)

	BeforeEach(func() {
		source = mock.NewRepositorySource()
		factory = &main.CommandFactory{RepositorySource: source}
		subject = expect.NoError(factory.RegisterRemoteRepositoriesCommand())
	})

	Describe("#Run", func() {
		It("registers remote repositories at the given URLs", func() {
			registered := testdata.NewURLs(
				"https://github.com/actions/checkout",
				"https://github.com/actions/setup-go",
			)

			subject.Run(registered)
			source.AddRemoteExpected("https://github.com/actions/checkout")
			source.AddRemoteExpected("https://github.com/actions/setup-go")
		})

		It("stops and returns an error, when failing to register a repository", func() {
			registered := testdata.NewURLs(
				"https://github.com/somebody/repo1",
				"https://github.com/somebody/repo2",
			)

			source.AddRemoteFails("https://github.com/somebody/repo1", "bang!")
			Expect(subject.Run(registered)).To(
				MatchError(ContainSubstring("failed to register https://github.com/somebody/repo1")))

			source.AddRemoteExpected("https://github.com/somebody/repo1")
			source.AddRemoteNotExpected("https://github.com/somebody/repo2")
		})

		It("returns nil, when there are no errors", func() {
			Expect(subject.Run(validUrls())).To(Succeed())
		})
	})
})

/* Test data */

func validUrls() []*url.URL {
	return testdata.NewURLs("https://github.com/actions/checkout")
}
