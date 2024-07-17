package svcfs

// Structure of the meta data file and its representation in JSON.
type metaRepoFile struct {
	MetaRepo metaRepoData `json:"meta_repo"`
	Version  string       `json:"version"`
}

type metaRepoData struct {
	RemoteRepositories []remoteRepositoryData `json:"remote_repositories"`
}

type remoteRepositoryData struct {
	Url string `json:"url"`
}
