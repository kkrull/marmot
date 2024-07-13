package corerepository

// Any number of Git repositories.
type Repositories interface {
	Names() []string
}

// Repositories backed by an array.
type RepositoriesArray struct {
	Repositories []Repository
}

func (array RepositoriesArray) Names() []string {
	names := make([]string, len(array.Repositories))
	for i, repository := range array.Repositories {
		names[i] = repository.Name
	}

	return names
}
