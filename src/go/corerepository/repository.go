package corerepository

import "net/url"

// One Git repository.
type Repository interface {
	LocalPath() string
	RemoteHref() string
	RemoteUrl() *url.URL
}

// Construct a repository with the given path.
func LocalRepository(localPath string) Repository {
	return &repository{localPath: localPath}
}

// Construct a repository with a URL to its remote host.
func RemoteRepository(remoteUrl *url.URL) Repository {
	return &repository{remoteUrl: remoteUrl}
}

// Construct a repository with something recognizable as a URL to its remote host.
func RemoteRepositoryS(remoteUrl string) (Repository, error) {
	if parsedUrl, parseErr := url.Parse(remoteUrl); parseErr != nil {
		return &repository{}, parseErr
	} else {
		return &repository{remoteUrl: parsedUrl}, nil
	}
}

// Basic Repository implementation.
type repository struct {
	localPath string
	remoteUrl *url.URL
}

func (repo repository) LocalPath() string {
	return repo.localPath
}

func (repo repository) RemoteHref() string {
	return repo.remoteUrl.String()
}

func (repo repository) RemoteUrl() *url.URL {
	return repo.remoteUrl
}
