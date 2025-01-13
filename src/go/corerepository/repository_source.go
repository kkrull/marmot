package corerepository

import "net/url"

// A source of Git repositories that a meta repo might care about.
//
// Repository sources are responsible for these invariants:
//   - Local repository paths are distinct, by string comparison.  Clients decide whether to
//     resolve relative paths or de-duplicate paths that resolve to the same filesystem entry.
//   - Remote repository URLs are distinct, comparing hrefs.
type RepositorySource interface {
	// Add repositories located at the specified paths, skipping known paths and exact duplicates.
	AddLocals(localPaths []string) error

	// Add repositories hosted at the specified URLs, skipping known remotes and duplicates.
	AddRemotes(hostUrls []*url.URL) error

	// List all known repositories on the local file system, including remotes cloned to known paths.
	ListLocal() (Repositories, error)

	// List all known repositories on remote hosts, including those that have not been cloned locally.
	ListRemote() (Repositories, error)
}
