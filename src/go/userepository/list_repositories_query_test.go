package userepository_test

import (
	repomock "github.com/kkrull/marmot/corerepositorymock"
	main "github.com/kkrull/marmot/mainfactory"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListRepositoriesQuery", func() {
	var (
		factory *main.CommandFactory
		source  *repomock.RepositorySource
	)

	Describe("#Run", func() {
		It("returns all repositories the source can list", func() {
			source = &repomock.RepositorySource{Names: []string{"one", "two"}}
			factory = &main.CommandFactory{RepositorySource: source}

			subject, factoryErr := factory.ListRepositoriesQuery()
			Expect(subject, factoryErr).NotTo(BeNil())

			repositories, runErr := subject.Run()
			Expect(repositories, runErr).NotTo(BeNil())
			Expect(repositories.Names()).To(ConsistOf("one", "two"))
		})
	})
})
