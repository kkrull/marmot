package corerepository

import "net/url"

// Construct a repository with the given path.
func LocalRepository(localPath string) Repository {
	return Repository{
		LocalPath: localPath,
		RemoteUrl: nil,
	}
}

// Construct a repository with a URL to its remote host.
func RemoteRepository(remoteUrl *url.URL) Repository {
	return Repository{
		LocalPath: "",
		RemoteUrl: remoteUrl,
	}
}

// Construct a repository with something recognizable as a URL to its remote host.
func RemoteRepositoryS(remoteUrl string) (Repository, error) {
	if parsedUrl, parseErr := url.Parse(remoteUrl); parseErr != nil {
		return Repository{}, parseErr
	} else {
		return Repository{
			LocalPath: "",
			RemoteUrl: parsedUrl,
		}, nil
	}
}

// One Git repository.
type Repository struct {
	LocalPath string
	RemoteUrl *url.URL
}
