package corerepository

// Any number of Git repositories.
type Repositories interface {
	Count() int
	Names() []string // TODO KDK: Make this more specific.  Is it the short name of the local repo or the remote one?
}

// Repositories backed by an array.
type RepositoriesArray struct {
	Repositories []Repository
}

func (array RepositoriesArray) Count() int {
	return len(array.Repositories)
}

func (array RepositoriesArray) Names() []string {
	names := make([]string, len(array.Repositories))
	for i, repository := range array.Repositories {
		names[i] = repository.Name
	}

	return names
}
