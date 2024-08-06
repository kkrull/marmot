package corerepository

import "net/url"

// A source of Git repositories that a meta repo might care about.
type RepositorySource interface {
	// TODO KDK: Consider adding #Add(Repository)
	AddLocal(localPath string) error
	AddRemote(hostUrl *url.URL) error
	ListLocal() (Repositories, error)
	ListRemote() (Repositories, error)
}
