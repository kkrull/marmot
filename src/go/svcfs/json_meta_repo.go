package svcfs

import (
	"fmt"
	"net/url"
	"slices"

	core "github.com/kkrull/marmot/corerepository"
)

func NewJsonMetaRepo(repositoryPath string) *JsonMetaRepo {
	return &JsonMetaRepo{
		localDataFile:  localDataFile(repositoryPath),
		sharedDataFile: sharedDataFile(repositoryPath),
	}
}

// A meta repo that stores meta data in JSON files in the specified directory.
type JsonMetaRepo struct {
	localDataFile  string
	sharedDataFile string
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
	return updateLocalFile(repo.localDataFile, func(rootObject *localRootObjectData) {
		rootObject.MetaRepo.AppendLocalRepository(localRepositoryData{Path: localPath})
	})
}

func (repo *JsonMetaRepo) ListLocal() (core.Repositories, error) {
	return queryLocalFile(
		repo.localDataFile,
		core.NoRepositories(),
		func(rootObject *localRootObjectData) (core.Repositories, error) {
			return rootObject.MetaRepo.MapLocalRepositories()
		})
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
	return updateFile(repo.sharedDataFile, func(rootObject *sharedRootObjectData) {
		rootObject.MetaRepo.AppendRemoteRepository(remoteRepositoryData{Url: hostUrl.String()})
	})
}

func (repo *JsonMetaRepo) ListRemote() (core.Repositories, error) {
	return querySharedFile(
		repo.sharedDataFile,
		core.NoRepositories(),
		func(rootObject *sharedRootObjectData) (core.Repositories, error) {
			return rootObject.MetaRepo.MapRemoteRepositories()
		})
}

/* I/O */

func queryLocalFile[V any](
	dataFile string,
	defaultValue V,
	queryData func(*localRootObjectData) (V, error),
) (V, error) {
	if rootObject, readErr := ReadLocalMetaRepoFile(dataFile); readErr != nil {
		return defaultValue, fmt.Errorf("failed to read file %s; %w", dataFile, readErr)
	} else if result, queryErr := queryData(rootObject); queryErr != nil {
		return defaultValue, fmt.Errorf("failed to query data; %w", queryErr)
	} else {
		return result, nil
	}
}

func querySharedFile[V any](
	dataFile string,
	defaultValue V,
	queryData func(*sharedRootObjectData) (V, error),
) (V, error) {
	if rootObject, readErr := ReadSharedMetaRepoFile(dataFile); readErr != nil {
		return defaultValue, fmt.Errorf("failed to read file %s; %w", dataFile, readErr)
	} else if result, queryErr := queryData(rootObject); queryErr != nil {
		return defaultValue, fmt.Errorf("failed to query data; %w", queryErr)
	} else {
		return result, nil
	}
}

type updateLocalDataFn = func(*localRootObjectData)

func updateLocalFile(dataFile string, updateData updateLocalDataFn) error {
	var rootObject *localRootObjectData
	rootObject, readErr := ReadLocalMetaRepoFile(dataFile)
	if readErr != nil {
		return fmt.Errorf("failed to read file %s; %w", dataFile, readErr)
	}

	updateData(rootObject)
	if writeErr := rootObject.WriteTo(dataFile); writeErr != nil {
		return fmt.Errorf("failed to write file %s; %w", dataFile, writeErr)
	} else {
		return nil
	}
}

type updateSharedDataFn = func(*sharedRootObjectData)

func updateFile(dataFile string, updateData updateSharedDataFn) error {
	var rootObject *sharedRootObjectData
	rootObject, readErr := ReadSharedMetaRepoFile(dataFile)
	if readErr != nil {
		return fmt.Errorf("failed to read file %s; %w", dataFile, readErr)
	}

	updateData(rootObject)
	if writeErr := rootObject.WriteTo(dataFile); writeErr != nil {
		return fmt.Errorf("failed to write file %s; %w", dataFile, writeErr)
	} else {
		return nil
	}
}
