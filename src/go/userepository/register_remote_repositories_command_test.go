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
			source.AddRemotesExpected(
				"https://github.com/actions/checkout",
				"https://github.com/actions/setup-go",
			)
		})

		It("returns no error, upon success", func() {
			Expect(subject.Run(validUrls())).To(Succeed())
		})

		It("returns an error, when adding repositories fails", func() {
			registered := testdata.NewURLs(
				"https://github.com/somebody/repo1",
				"https://github.com/somebody/repo2",
			)

			source.AddRemotesFails("bang!")
			Expect(subject.Run(registered)).To(
				MatchError("failed to register remote repositories; bang!"))
		})
	})
})

/* Test data */

func validUrls() []*url.URL {
	return testdata.NewURLs("https://github.com/actions/checkout")
}
