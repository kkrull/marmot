package userepository_test

import (
	core "github.com/kkrull/marmot/corerepository"
	main "github.com/kkrull/marmot/mainfactory"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListRepositoriesQuery", func() {
	var (
		factory *main.CommandFactory
		source  *MockRepositorySource
	)

	Describe("#Run", func() {
		It("returns all repositories included the source can list", func() {
			source = &MockRepositorySource{Names: []string{"one", "two"}}
			factory = &main.CommandFactory{RepositorySource: source}

			subject, factoryErr := factory.ListRepositoriesQuery()
			Expect(subject, factoryErr).NotTo(BeNil())

			repositories, runErr := subject.Run()
			Expect(repositories, runErr).NotTo(BeNil())
			Expect(repositories.Names()).To(ConsistOf("one", "two"))
		})
	})
})

type MockRepositorySource struct {
	Names []string
}

func (source *MockRepositorySource) List() (core.Repositories, error) {
	repositories := make([]core.Repository, 0, len(source.Names))
	for _, name := range source.Names {
		repositories = append(repositories, core.Repository{Name: name})
	}

	return &core.RepositoriesArray{Repositories: repositories}, nil
}
