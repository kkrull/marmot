package svcfs

// Construct a meta repo file with only minimal information; e.g. no Git repositories.
func EmptyMetaRepoFile(version string) *rootObjectData {
	return &rootObjectData{
		MetaRepo: metaRepoData{
			RemoteRepositories: make([]remoteRepositoryData, 0),
		},
		Version: version,
	}
}

// Top-level object in the meta data file and its representation in JSON.
type rootObjectData struct {
	MetaRepo metaRepoData `json:"meta_repo"`
	Version  string       `json:"version"`
}

type metaRepoData struct {
	RemoteRepositories []remoteRepositoryData `json:"remote_repositories"`
}

type remoteRepositoryData struct {
	Url string `json:"url"`
}
