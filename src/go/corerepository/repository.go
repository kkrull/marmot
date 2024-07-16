package corerepository

import "net/url"

// One Git repository.
type Repository struct {
	RemoteUrl *url.URL // TODO KDK: Add RemoteRepository constructor
}
