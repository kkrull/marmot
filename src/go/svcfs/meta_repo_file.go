package svcfs

/* Constructors */

// Construct a meta repo file with only minimal information; e.g. no Git repositories.
func InitMetaRepoData(version string) *rootObjectData {
	return &rootObjectData{
		MetaRepo: &metaRepoData{
			RemoteRepositories: make([]remoteRepositoryData, 0),
		},
		Version: version,
	}
}

/* Structure */

// Top-level object in the meta data file and its representation in JSON.
type rootObjectData struct {
	MetaRepo *metaRepoData `json:"meta_repo"`
	Version  string        `json:"version"`
}

type metaRepoData struct {
	RemoteRepositories []remoteRepositoryData `json:"remote_repositories"`
}

type remoteRepositoryData struct {
	Url string `json:"url"`
}

/* Updates */

func (metaRepo *metaRepoData) AppendRemoteRepository(remoteRepository remoteRepositoryData) {
	metaRepo.RemoteRepositories = append(metaRepo.RemoteRepositories, remoteRepository)
}
