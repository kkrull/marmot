package corerepository

import "net/url"

// A source of Git repositories that a meta repo might care about.
type RepositorySource interface {
	AddRemote(hostUrl *url.URL) error
	ListRemote() (Repositories, error)
}
