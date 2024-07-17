package svcfs

import (
	"errors"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"

	core "github.com/kkrull/marmot/corerepository"
)

// A meta repo that stores meta data in JSON files in the specified directory.
func NewJsonMetaDataRepo(repositoryPath string) *JsonMetaDataRepo {
	metaDataDir := filepath.Join(repositoryPath, ".marmot")
	return &JsonMetaDataRepo{
		repositoryDir: repositoryPath,
		metaDataFile:  filepath.Join(metaDataDir, "meta-repo.json"),
	}
}

type JsonMetaDataRepo struct {
	// metaDataDir   string
	metaDataFile  string
	repositoryDir string
}

/* MetaDataAdmin */

func (*JsonMetaDataRepo) InitP(repositoryDir string) error {
	_, statErr := os.Stat(repositoryDir)
	if errors.Is(statErr, fs.ErrNotExist) {
		return initDirectory(repositoryDir)
	} else if statErr != nil {
		return fmt.Errorf("failed to check for existing meta repo %s; %w", repositoryDir, statErr)
	} else {
		return fmt.Errorf("path already exists: %s", repositoryDir)
	}
}

func initDirectory(repositoryDir string) error {
	metaDataDir := filepath.Join(repositoryDir, ".marmot")
	metaDataFile := filepath.Join(metaDataDir, "meta-repo.json")

	emptyFile := EmptyMetaRepoFile("0.0.1")
	if dirErr := os.MkdirAll(metaDataDir, fs.ModePerm); dirErr != nil {
		return fmt.Errorf("failed to make directory %s; %w", metaDataDir, dirErr)
	} else if writeErr := emptyFile.WriteTo(metaDataFile); writeErr != nil {
		return fmt.Errorf("failed to write file %s; %w", metaDataFile, writeErr)
	} else {
		return nil
	}
}

/* RepositorySource */

func (repo *JsonMetaDataRepo) AddRemote(hostUrl *url.URL) error {
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

func (repo *JsonMetaDataRepo) ListRemote() (core.Repositories, error) {
	if rootObject, readErr := ReadMetaRepoFile(repo.metaDataFile); readErr != nil {
		return nil, fmt.Errorf("failed to read file %s; %w", repo.metaDataFile, readErr)
	} else if repositories, mapErr := rootObject.MetaRepo.MapRemoteRepositories(); mapErr != nil {
		return nil, fmt.Errorf("failed to map to core model; %w", mapErr)
	} else {
		return repositories, nil
	}
}
