package corerepository

import "net/url"

// A source of Git repositories that a meta repo might care about.
type RepositorySource interface {
	// TODO KDK: Consider adding #Add(Repository)

	// Add a repository located at the specified path on the filesystem.
	AddLocal(localPath string) error

	// Add a repository hosted on the specified URL.
	AddRemote(hostUrl *url.URL) error

	// List all known repositories on the local file system, including remotes cloned to known paths.
	ListLocal() (Repositories, error)

	// List all known repositories on remote hosts, including those that have not been cloned locally.
	ListRemote() (Repositories, error)
}
