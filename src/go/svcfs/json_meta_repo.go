package svcfs

import (
	"fmt"
	"net/url"

	core "github.com/kkrull/marmot/corerepository"
)

func NewJsonMetaRepo(repositoryPath string) *JsonMetaRepo {
	return &JsonMetaRepo{metaDataFile: metaDataFile(repositoryPath)}
}

// A meta repo that stores meta data in JSON files in the specified directory.
type JsonMetaRepo struct {
	metaDataFile string
}

/* RepositorySource */

func (repo *JsonMetaRepo) AddRemote(hostUrl *url.URL) error {
	var rootObject *rootObjectData
	rootObject, readErr := ReadMetaRepoFile(repo.metaDataFile)
	if readErr != nil {
		return fmt.Errorf("failed to read file %s; %w", repo.metaDataFile, readErr)
	}

	rootObject.MetaRepo.AppendRemoteRepository(remoteRepositoryData{Url: hostUrl.String()})
	if writeErr := rootObject.WriteTo(repo.metaDataFile); writeErr != nil {
		return fmt.Errorf("failed to write file %s; %w", repo.metaDataFile, writeErr)
	} else {
		return nil
	}
}

func (repo *JsonMetaRepo) ListRemote() (core.Repositories, error) {
	if rootObject, readErr := ReadMetaRepoFile(repo.metaDataFile); readErr != nil {
		return nil, fmt.Errorf("failed to read file %s; %w", repo.metaDataFile, readErr)
	} else if repositories, mapErr := rootObject.MetaRepo.MapRemoteRepositories(); mapErr != nil {
		return nil, fmt.Errorf("failed to map to core model; %w", mapErr)
	} else {
		return repositories, nil
	}
}
