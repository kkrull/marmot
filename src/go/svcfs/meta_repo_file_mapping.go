package svcfs

import (
	"fmt"

	"github.com/kkrull/marmot/corerepository"
)

func (metaRepo *metaRepoData) MapRemoteRepositories() (corerepository.Repositories, error) {
	repositories := make([]corerepository.Repository, len(metaRepo.RemoteRepositories))
	for i, remoteRepository := range metaRepo.RemoteRepositories {
		if repository, mapErr := remoteRepository.ToCoreRepository(); mapErr != nil {
			return nil, mapErr
		} else {
			repositories[i] = repository
		}
	}

	return &corerepository.RepositoriesArray{
		Repositories: repositories,
	}, nil
}

func (remoteRepo *remoteRepositoryData) ToCoreRepository() (corerepository.Repository, error) {
	if repository, err := corerepository.RemoteRepositoryS(remoteRepo.Url); err != nil {
		return corerepository.Repository{}, fmt.Errorf("failed to parse %s; %w", remoteRepo.Url, err)
	} else {
		return repository, nil
	}
}
