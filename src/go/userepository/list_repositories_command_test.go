package userepository_test

import (
	core "github.com/kkrull/marmot/corerepository"
	main "github.com/kkrull/marmot/mainfactory"
	use "github.com/kkrull/marmot/userepository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListRepositoriesCommand", func() {
	var (
		factory *main.CommandFactory
		source  *MockRepositorySource
		subject *use.ListRepositoriesCommand
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

func (source *MockRepositorySource) List() (core.Repositories, error) {
	repositories := make([]core.Repository, len(source.Names))
	for _, name := range source.Names {
		repositories = append(repositories, core.Repository{Name: name})
	}

	return &RepositoriesArray{Repositories: repositories}, nil
}

type RepositoriesArray struct {
	Repositories []core.Repository
}

func (array RepositoriesArray) Names() []string {
	names := make([]string, len(array.Repositories))
	for _, repository := range array.Repositories {
		names = append(names, repository.Name)
	}

	return names
}
