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
		source  *MockRepositorySource
		subject *repository.ListRepositoriesCommand
	)

	Describe("#Run", func() {
		It("returns an empty result set, given a meta repo with no repositories", func() {
			source = &MockRepositorySource{Names: make([]string, 0)}
			factory = &main.CommandFactory{RepositorySource: source}

			subject, _ = factory.ListRepositoriesCommand()
			repositories, _ := subject.Run()
			Expect(repositories.Names()).To(BeEmpty())
		})

		It("returns something interesting", Pending, func() {
		})
	})
})

type MockRepositorySource struct {
	Names []string
}

func (source *MockRepositorySource) List() (repository.Repositories, error) {
	repositories := make([]repository.Repository, len(source.Names))
	for _, name := range source.Names {
		repositories = append(repositories, repository.Repository{Name: name})
	}

	return &RepositoriesArray{Repositories: repositories}, nil
}

type RepositoriesArray struct {
	Repositories []repository.Repository
}

func (array RepositoriesArray) Names() []string {
	names := make([]string, len(array.Repositories))
	for _, repository := range array.Repositories {
		names = append(names, repository.Name)
	}

	return names
}
