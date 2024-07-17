package corerepository

import "net/url"

// Construct a repository with a URL to its remote host.
func RemoteRepository(remoteUrl *url.URL) Repository {
	return Repository{RemoteUrl: remoteUrl}
}

// Construct a repository with something recognizable as a URL to its remote host.
func RemoteRepositoryS(remoteUrl string) (Repository, error) {
	if parsedUrl, parseErr := url.Parse(remoteUrl); parseErr != nil {
		return Repository{}, parseErr
	} else {
		return Repository{RemoteUrl: parsedUrl}, nil
	}
}

// One Git repository.
type Repository struct {
	RemoteUrl *url.URL // TODO KDK: Add RemoteRepository constructor
}
