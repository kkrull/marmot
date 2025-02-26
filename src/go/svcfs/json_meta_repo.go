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
	return updateFile(
		repo.localDataFile,
		ReadLocalMetaRepoFile,
		func(rootObject *localRootObjectData) {
			rootObject.MetaRepo.AppendLocalRepository(localRepositoryData{Path: localPath})
		},
		func(rootObject *localRootObjectData, dataFile string) error {
			return rootObject.WriteTo(dataFile)
		},
	)
}

func (repo *JsonMetaRepo) ListLocal() (core.Repositories, error) {
	return queryFile(
		repo.localDataFile,
		core.NoRepositories(),
		ReadLocalMetaRepoFile,
		func(rootObject *localRootObjectData) (core.Repositories, error) {
			return rootObject.MetaRepo.MapLocalRepositories()
		},
	)
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
	return updateFile(
		repo.sharedDataFile,
		ReadSharedMetaRepoFile,
		func(rootObject *sharedRootObjectData) {
			rootObject.MetaRepo.AppendRemoteRepository(remoteRepositoryData{Url: hostUrl.String()})
		},
		func(rootObject *sharedRootObjectData, dataFile string) error {
			return rootObject.WriteTo(dataFile)
		},
	)
}

func (repo *JsonMetaRepo) ListRemote() (core.Repositories, error) {
	return queryFile(
		repo.sharedDataFile,
		core.NoRepositories(),
		ReadSharedMetaRepoFile,
		func(rootObject *sharedRootObjectData) (core.Repositories, error) {
			return rootObject.MetaRepo.MapRemoteRepositories()
		})
}

/* I/O */

func queryFile[RootObject any, Result any](
	dataFile string,
	defaultValue Result,
	fetchData func(string) (RootObject, error),
	queryData func(RootObject) (Result, error),
) (Result, error) {
	if rootObject, readErr := fetchData(dataFile); readErr != nil {
		return defaultValue, fmt.Errorf("failed to read file %s; %w", dataFile, readErr)
	} else if result, queryErr := queryData(rootObject); queryErr != nil {
		return defaultValue, fmt.Errorf("failed to query data; %w", queryErr)
	} else {
		return result, nil
	}
}

func updateFile[RootObject any](
	dataFile string,
	fetchData func(string) (RootObject, error),
	updateData func(RootObject),
	writeData func(RootObject, string) error,
) error {
	var rootObject RootObject
	rootObject, readErr := fetchData(dataFile)
	if readErr != nil {
		return fmt.Errorf("failed to read file %s; %w", dataFile, readErr)
	}

	updateData(rootObject)
	if writeErr := writeData(rootObject, dataFile); writeErr != nil {
		return fmt.Errorf("failed to write file %s; %w", dataFile, writeErr)
	} else {
		return nil
	}
}
