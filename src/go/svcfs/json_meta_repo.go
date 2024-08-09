package svcfs

import (
	"fmt"
	"net/url"
	"slices"

	core "github.com/kkrull/marmot/corerepository"
)

func NewJsonMetaRepo(repositoryPath string) *JsonMetaRepo {
	return &JsonMetaRepo{metaDataFile: metaDataFile(repositoryPath)}
}

// A meta repo that stores meta data in JSON files in the specified directory.
type JsonMetaRepo struct {
	metaDataFile string
}

/* Local repositories */

func (repo *JsonMetaRepo) AddLocals(localPaths []string) error {
	knownPaths := make([]string, 0)
	if alreadyRegistered, listErr := repo.ListLocal(); listErr == nil {
		knownPaths = append(knownPaths, alreadyRegistered.LocalPaths()...)
	}

	for _, localPath := range localPaths {
		if isDuplicate := slices.Contains(knownPaths, localPath); isDuplicate {
			continue
		} else if err := repo.addLocal(localPath); err != nil {
			return err
		} else {
			knownPaths = append(knownPaths, localPath)
		}
	}

	return nil
}

func (repo *JsonMetaRepo) addLocal(localPath string) error {
	return repo.updateFile(func(rootObject *rootObjectData) {
		rootObject.MetaRepo.AppendLocalRepository(localRepositoryData{Path: localPath})
	})
}

func (repo *JsonMetaRepo) ListLocal() (core.Repositories, error) {
	if rootObject, readErr := ReadMetaRepoFile(repo.metaDataFile); readErr != nil {
		return nil, fmt.Errorf("failed to read file %s; %w", repo.metaDataFile, readErr)
	} else if repositories, mapErr := rootObject.MetaRepo.MapLocalRepositories(); mapErr != nil {
		return nil, fmt.Errorf("failed to map to core model; %w", mapErr)
	} else {
		return repositories, nil
	}
}

/* Remote repositories */

func (repo *JsonMetaRepo) AddRemotes(hostUrls []*url.URL) error {
	knownHrefs := make([]string, 0)
	if alreadyRegistered, listErr := repo.ListRemote(); listErr == nil {
		knownHrefs = append(knownHrefs, alreadyRegistered.RemoteHrefs()...)
	}

	for _, hostUrl := range hostUrls {
		if isDuplicate := slices.Contains(knownHrefs, hostUrl.String()); isDuplicate {
			continue
		} else if err := repo.addRemote(hostUrl); err != nil {
			return err
		} else {
			knownHrefs = append(knownHrefs, hostUrl.String())
		}
	}

	return nil
}

func (repo *JsonMetaRepo) addRemote(hostUrl *url.URL) error {
	return repo.updateFile(func(rootObject *rootObjectData) {
		rootObject.MetaRepo.AppendRemoteRepository(remoteRepositoryData{Url: hostUrl.String()})
	})
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

/* I/O */

type updateDataFn = func(*rootObjectData)

func (repo *JsonMetaRepo) updateFile(updateData updateDataFn) error {
	var rootObject *rootObjectData
	rootObject, readErr := ReadMetaRepoFile(repo.metaDataFile)
	if readErr != nil {
		return fmt.Errorf("failed to read file %s; %w", repo.metaDataFile, readErr)
	}

	updateData(rootObject)
	if writeErr := rootObject.WriteTo(repo.metaDataFile); writeErr != nil {
		return fmt.Errorf("failed to write file %s; %w", repo.metaDataFile, writeErr)
	} else {
		return nil
	}
}
