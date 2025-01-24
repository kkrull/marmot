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

var _ = Describe("RegisterRemoteRepositoriesAction", func() {
	var (
		subject *userepository.RegisterRemoteRepositoriesAction
		source  *mock.RepositorySource
	)

	validUrls := func() []*url.URL {
		return testdata.NewURLs("https://github.com/actions/checkout")
	}

	BeforeEach(func() {
		source = mock.NewRepositorySource()
		factory := use.NewActionFactory().WithRemoteRepositorySource(source)
		subject = expect.NoError(factory.NewRegisterRemoteRepositories())
	})

	Describe("#Run", func() {
		It("registers remote repositories at the given URLs", func() {
			given := testdata.NewURLs(
				"https://github.com/actions/checkout",
				"https://github.com/actions/setup-go",
			)

			subject.Run(given)
			source.AddRemotesExpected(
				"https://github.com/actions/checkout",
				"https://github.com/actions/setup-go",
			)
		})

		It("returns no error, upon success", func() {
			Expect(subject.Run(validUrls())).To(Succeed())
		})

		It("returns an error, when adding repositories fails", func() {
			source.AddRemotesFails("bang!")
			Expect(subject.Run(validUrls())).To(
				MatchError("failed to register remote repositories; bang!"))
		})
	})
})
