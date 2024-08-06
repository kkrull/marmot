package svcfs

import (
	"fmt"

	"github.com/kkrull/marmot/corerepository"
)

/* Local repositories */

func (metaRepo *metaRepoData) MapLocalRepositories() (corerepository.Repositories, error) {
	repositories := make([]corerepository.Repository, len(metaRepo.LocalRepositories))
	for i, localRepositoryData := range metaRepo.LocalRepositories {
		if repository, mapErr := localRepositoryData.ToCoreRepository(); mapErr != nil {
			return nil, mapErr
		} else {
			repositories[i] = repository
		}
	}

	return corerepository.SomeRepositories(repositories), nil
}

func (localRepo *localRepositoryData) ToCoreRepository() (corerepository.Repository, error) {
	return corerepository.LocalRepository(localRepo.Path), nil
}

/* Remote repositories */

func (metaRepo *metaRepoData) MapRemoteRepositories() (corerepository.Repositories, error) {
	repositories := make([]corerepository.Repository, len(metaRepo.RemoteRepositories))
	for i, remoteRepositoryData := range metaRepo.RemoteRepositories {
		if repository, mapErr := remoteRepositoryData.ToCoreRepository(); mapErr != nil {
			return nil, mapErr
		} else {
			repositories[i] = repository
		}
	}

	return corerepository.SomeRepositories(repositories), nil
}

func (remoteRepo *remoteRepositoryData) ToCoreRepository() (corerepository.Repository, error) {
	if repository, err := corerepository.RemoteRepositoryS(remoteRepo.Url); err != nil {
		return nil, fmt.Errorf("failed to parse %s; %w", remoteRepo.Url, err)
	} else {
		return repository, nil
	}
}
