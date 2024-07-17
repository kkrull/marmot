package userepository_test

import (
	repomock "github.com/kkrull/marmot/corerepositorymock"
	main "github.com/kkrull/marmot/mainfactory"
	testdata "github.com/kkrull/marmot/testsupportdata"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListRemoteRepositoriesQuery", func() {
	var (
		factory *main.CommandQueryFactory
		source  *repomock.RepositorySource
	)

	Describe("#Run", func() {
		It("returns all repositories the source can list", func() {
			remoteUrls := testdata.NewURLs(
				"https://github.com/actions/checkout",
				"https://github.com/actions/setup-up",
			)
			source = &repomock.RepositorySource{ListRemoteUrls: remoteUrls}
			factory = &main.CommandQueryFactory{RepositorySource: source}

			subject, factoryErr := factory.ListRepositoriesQuery()
			Expect(subject, factoryErr).NotTo(BeNil())

			repositories, runErr := subject.Run()
			Expect(repositories, runErr).NotTo(BeNil())
			Expect(repositories.RemoteHrefs()).To(ConsistOf(
				"https://github.com/actions/checkout",
				"https://github.com/actions/setup-up",
			))
		})
	})
})
