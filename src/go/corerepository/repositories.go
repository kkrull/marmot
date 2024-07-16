package corerepository

import "net/url"

// Any number of Git repositories.
type Repositories interface {
	Count() int
	Names() []string // TODO KDK: Remove
	RemoteUrls() []url.URL
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

func (array RepositoriesArray) RemoteUrls() []url.URL {
	remoteUrls := make([]url.URL, len(array.Repositories))
	for i, repository := range array.Repositories {
		remoteUrls[i] = repository.RemoteUrl
	}

	return remoteUrls
}
