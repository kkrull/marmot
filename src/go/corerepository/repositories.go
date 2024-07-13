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
	names := make([]string, 0, len(array.Repositories))
	for _, repository := range array.Repositories {
		names = append(names, repository.Name)
	}

	return names
}
