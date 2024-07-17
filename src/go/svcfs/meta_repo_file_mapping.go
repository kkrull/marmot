package svcfs

import (
	"fmt"

	"github.com/kkrull/marmot/corerepository"
)

func (root *rootObjectData) ToCoreRepositories() (corerepository.Repositories, error) {
	repositories := make([]corerepository.Repository, len(root.MetaRepo.RemoteRepositories))
	for i, remoteRepositoryData := range root.MetaRepo.RemoteRepositories {
		if remoteRepository, parseErr := corerepository.RemoteRepositoryS(remoteRepositoryData.Url); parseErr != nil {
			return nil, fmt.Errorf("failed to parse %s; %w", remoteRepositoryData.Url, parseErr)
		} else {
			repositories[i] = remoteRepository
		}
	}

	return &corerepository.RepositoriesArray{
		Repositories: repositories,
	}, nil
}
