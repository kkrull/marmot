package corerepository

import "net/url"

// One Git repository.
type Repository struct {
	Name      string //TODO KDK: Replace with more specific things like shortPath, shortUrl, etc..
	RemoteUrl url.URL
}
