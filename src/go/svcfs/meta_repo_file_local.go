package svcfs

/* Constructors */

// Construct a meta repo file with only minimal information; e.g. no Git repositories.
func InitLocalMetaRepoData(version string) *localRootObjectData {
	return &localRootObjectData{
		MetaRepo: &localMetaRepoData{
			LocalRepositories: make([]localRepositoryData, 0),
		},
		Version: version,
	}
}

/* Structure */

// Top-level object in the meta data file and its representation in JSON.
type localRootObjectData struct {
	MetaRepo *localMetaRepoData `json:"meta_repo"`
	Version  string             `json:"version"`
}

type localMetaRepoData struct {
	LocalRepositories []localRepositoryData `json:"local_repositories"`
}

type localRepositoryData struct {
	Path string `json:"path"`
}

/* Updates */

func (metaRepo *localMetaRepoData) AppendLocalRepository(localRepository localRepositoryData) {
	metaRepo.LocalRepositories = append(metaRepo.LocalRepositories, localRepository)
}
