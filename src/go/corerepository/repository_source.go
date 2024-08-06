package corerepository

import "net/url"

// A source of Git repositories that a meta repo might care about.
type RepositorySource interface {
	AddLocal(localPath string) error
	AddRemote(hostUrl *url.URL) error
	ListRemote() (Repositories, error)
}
