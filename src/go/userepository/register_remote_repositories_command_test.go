package userepository_test

import (
	"net/url"

	mock "github.com/kkrull/marmot/corerepositorymock"
	testdata "github.com/kkrull/marmot/testsupportdata"
	expect "github.com/kkrull/marmot/testsupportexpect"
	"github.com/kkrull/marmot/use"
	"github.com/kkrull/marmot/userepository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegisterRemoteRepositoriesCommand", func() {
	var (
		factory use.CommandFactory
		source  *mock.RepositorySource
		subject *userepository.RegisterRemoteRepositoriesCommand
	)

	BeforeEach(func() {
		source = mock.NewRepositorySource()
		factory = use.NewCommandFactory().WithRepositorySource(source)
		subject = expect.NoError(factory.NewRegisterRemoteRepositories())
	})

	Describe("#Run", func() {
		It("registers remote repositories at the given URLs", func() {
			registered := testdata.NewURLs(
				"https://github.com/actions/checkout",
				"https://github.com/actions/setup-go",
			)

			subject.Run(registered)
			source.AddRemoteExpected(
				"https://github.com/actions/checkout",
				"https://github.com/actions/setup-go",
			)
		})

		It("returns no error, upon success", func() {
			Expect(subject.Run(validUrls())).To(Succeed())
		})

		//could be unreachable behind corporate firewall for now, but maybe that doesn't matter?
		It("rejects URLs that are not Git repositories?", Pending)

		//duplicating URLs doesn't make sense (so it's a business rule), but the repository source will do this atomically anyway
		It("skips hosts that have already been added?", Pending)

		//no news is good news on a CLI, but a web UI might show some sort of confirmation status.
		It("reports the URLs that were actually added?", Pending)

		It("stops and returns an error, after when adding a repository fails", func() {
			registered := testdata.NewURLs(
				"https://github.com/somebody/repo1",
				"https://github.com/somebody/repo2",
			)

			source.AddRemoteFails("https://github.com/somebody/repo1", "bang!")
			Expect(subject.Run(registered)).To(
				MatchError(ContainSubstring("failed to register https://github.com/somebody/repo1")))

			source.AddRemoteExpected("https://github.com/somebody/repo1")
		})
	})
})

/* Test data */

func validUrls() []*url.URL {
	return testdata.NewURLs("https://github.com/actions/checkout")
}
