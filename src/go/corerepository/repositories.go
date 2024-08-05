package corerepository

import "net/url"

// Any number of Git repositories.
type Repositories interface {
	// How many repositories are in this collection
	Count() int

	// Paths to each repository on the local file system
	LocalPaths() []string

	// String versions of each remote's URL
	RemoteHrefs() []string

	// URL to each remote repository
	RemoteUrls() []*url.URL
}

// Repositories backed by an array.
type RepositoriesArray struct {
	Repositories []Repository
}

func (array RepositoriesArray) Count() int {
	return len(array.Repositories)
}

func (array RepositoriesArray) LocalPaths() []string {
	localPaths := make([]string, len(array.Repositories))
	for i, repository := range array.Repositories {
		localPaths[i] = repository.LocalPath
	}

	return localPaths
}

func (array RepositoriesArray) RemoteHrefs() []string {
	remoteHrefs := make([]string, len(array.Repositories))
	for i, repository := range array.Repositories {
		remoteHrefs[i] = repository.RemoteUrl.String()
	}

	return remoteHrefs
}

func (array RepositoriesArray) RemoteUrls() []*url.URL {
	remoteUrls := make([]*url.URL, len(array.Repositories))
	for i, repository := range array.Repositories {
		remoteUrls[i] = repository.RemoteUrl
	}

	return remoteUrls
}
