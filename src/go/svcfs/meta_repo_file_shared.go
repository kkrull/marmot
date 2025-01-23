package svcfs

/* Constructors */

// Construct a meta repo file with only minimal information; e.g. no Git repositories.
func InitSharedMetaRepoData(version string) *sharedRootObjectData {
	return &sharedRootObjectData{
		MetaRepo: &sharedMetaRepoData{
			RemoteRepositories: make([]remoteRepositoryData, 0),
		},
		Version: version,
	}
}

/* Structure */

// Top-level object in the meta data file and its representation in JSON.
type sharedRootObjectData struct {
	MetaRepo *sharedMetaRepoData `json:"meta_repo"`
	Version  string              `json:"version"`
}

type sharedMetaRepoData struct {
	RemoteRepositories []remoteRepositoryData `json:"remote_repositories"`
}

type remoteRepositoryData struct {
	Url string `json:"url"`
}

/* Updates */

func (metaRepo *sharedMetaRepoData) AppendRemoteRepository(remoteRepository remoteRepositoryData) {
	metaRepo.RemoteRepositories = append(metaRepo.RemoteRepositories, remoteRepository)
}
