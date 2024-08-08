package corerepository

import "net/url"

// A source of Git repositories that a meta repo might care about.
//   - Remote repository URLs are distinct.
type RepositorySource interface {
	// TODO KDK: Consider adding #Add(Repository)

	// Add a repository located at the specified path on the filesystem.
	AddLocal(localPath string) error

	// Add repositories hosted at the specified URLs, skipping known remotes and duplicates.
	AddRemotes(hostUrls []*url.URL) error

	// List all known repositories on the local file system, including remotes cloned to known paths.
	ListLocal() (Repositories, error)

	// List all known repositories on remote hosts, including those that have not been cloned locally.
	ListRemote() (Repositories, error)
}
