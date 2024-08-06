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

// Construct a container containing the given repositories.
func SomeRepositories(repositories []Repository) Repositories {
	return &repositoriesArray{Repositories: repositories}
}

// Construct a container containing no repositories of any kind.
func NoRepositories() Repositories {
	return &repositoriesArray{
		Repositories: make([]Repository, 0),
	}
}

// Repositories backed by an array.
type repositoriesArray struct {
	Repositories []Repository
}

func (array repositoriesArray) Count() int {
	return len(array.Repositories)
}

func (array repositoriesArray) LocalPaths() []string {
	localPaths := make([]string, len(array.Repositories))
	for i, repository := range array.Repositories {
		localPaths[i] = repository.LocalPath
	}

	return localPaths
}

func (array repositoriesArray) RemoteHrefs() []string {
	remoteHrefs := make([]string, len(array.Repositories))
	for i, repository := range array.Repositories {
		remoteHrefs[i] = repository.RemoteUrl.String()
	}

	return remoteHrefs
}

func (array repositoriesArray) RemoteUrls() []*url.URL {
	remoteUrls := make([]*url.URL, len(array.Repositories))
	for i, repository := range array.Repositories {
		remoteUrls[i] = repository.RemoteUrl
	}

	return remoteUrls
}
