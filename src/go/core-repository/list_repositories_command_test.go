package core_repository_test

import (
	repository "github.com/kkrull/marmot/core-repository"
	main "github.com/kkrull/marmot/main-factory"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListRepositoriesCommand", func() {
	var (
		factory *main.CommandFactory
		subject *repository.ListRepositoriesCommand
	)

	Describe("#Run", func () {
		It("returns an empty result set, given a meta repo with no repositories", Pending, func () {
			factory = &main.CommandFactory{}
			subject, _ = factory.ListRepositoriesCommand() //TODO KDK: Provide mock data source, first
			repositories, _ := subject.Run()
			Expect(repositories.Names()).To(BeEmpty())
		})
	})
})
