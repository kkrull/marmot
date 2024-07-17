package svcfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"

	"github.com/kkrull/marmot/corerepository"
)

// A meta repo that stores meta data in JSON files on the file system.
func NewJsonMetaDataRepo(repositoryPath string) *JsonMetaDataRepo {
	metaDataDir := filepath.Join(repositoryPath, ".marmot")
	return &JsonMetaDataRepo{
		repositoryDir: repositoryPath,
		metaDataDir:   metaDataDir,
		metaDataFile:  filepath.Join(metaDataDir, "meta-repo.json"),
	}
}

type JsonMetaDataRepo struct {
	metaDataDir   string
	metaDataFile  string
	repositoryDir string
}

/* MetaDataAdmin */

func (meta *JsonMetaDataRepo) Init() error {
	_, statErr := os.Stat(meta.repositoryDir)
	if errors.Is(statErr, fs.ErrNotExist) {
		return meta.initDirectory()
	} else if statErr != nil {
		return fmt.Errorf("failed to check for existing meta repo %s; %w", meta.repositoryDir, statErr)
	} else {
		return fmt.Errorf("path already exists: %s", meta.repositoryDir)
	}
}

func (meta *JsonMetaDataRepo) initDirectory() error {
	if dirErr := os.MkdirAll(meta.metaDataDir, fs.ModePerm); dirErr != nil {
		return fmt.Errorf("failed to make directory %s; %w", meta.metaDataDir, dirErr)
	} else if metaDataFd, fileErr := os.Create(meta.metaDataFile); fileErr != nil {
		return fmt.Errorf("failed to create file %s; %w", meta.metaDataFile, fileErr)
	} else {
		defer metaDataFd.Close()
		return meta.writeToFile(metaDataFd, make([]remoteRepositoryData, 0))
	}
}

/* RepositorySource */

func (metaRepo *JsonMetaDataRepo) List() (corerepository.Repositories, error) {
	var decoder *json.Decoder
	if metaDataFd, openErr := os.Open(metaRepo.metaDataFile); openErr != nil {
		return nil, fmt.Errorf("failed to open file %s; %w", metaRepo.metaDataFile, openErr)
	} else {
		decoder = json.NewDecoder(metaDataFd)
		// defer metaDataFd.Close() //TODO KDK: Test and restore
	}

	var content metaRepoFile
	if decodeErr := decoder.Decode(&content); decodeErr != nil {
		return nil, fmt.Errorf("failed to decode %s; %w", metaRepo.metaDataFile, decodeErr)
	}

	repositories := make([]corerepository.Repository, len(content.MetaRepo.RemoteRepositories))
	for i, remoteRepositoryData := range content.MetaRepo.RemoteRepositories {
		if remoteRepository, parseErr := corerepository.RemoteRepositoryS(remoteRepositoryData.Url); parseErr != nil {
			return nil, fmt.Errorf("failed to parse %s from %s; %w", remoteRepositoryData.Url, metaRepo.metaDataFile, parseErr)
		} else {
			repositories[i] = remoteRepository
		}
	}

	return &corerepository.RepositoriesArray{
		Repositories: repositories,
	}, nil
}

func (metaRepo *JsonMetaDataRepo) RegisterRemote(hostUrl *url.URL) error {
	var encoder *json.Encoder
	if metaDataFd, openErr := os.OpenFile(metaRepo.metaDataFile, os.O_WRONLY, os.ModePerm); openErr != nil {
		return fmt.Errorf("failed to open file %s; %w", metaRepo.metaDataFile, openErr)
	} else {
		encoder = json.NewEncoder(metaDataFd)
	}

	metaRepoFile := &metaRepoFile{
		MetaRepo: metaRepoData{
			RemoteRepositories: []remoteRepositoryData{
				{Url: hostUrl.String()},
			},
		},
		Version: "0.1",
	}

	if encodeErr := encoder.Encode(metaRepoFile); encodeErr != nil {
		return fmt.Errorf("failed to encode content; %w", encodeErr)
	} else {
		return nil
	}
}

/* JSON encoding */

func (meta *JsonMetaDataRepo) writeToFile(file *os.File, remoteRepositories []remoteRepositoryData) error {
	content := &metaRepoFile{
		MetaRepo: metaRepoData{
			RemoteRepositories: remoteRepositories,
		},
		Version: "0.1",
	}

	encoder := *json.NewEncoder(file)
	if encodeErr := encoder.Encode(content); encodeErr != nil {
		return fmt.Errorf("failed to encode JSON data; %w", encodeErr)
	} else {
		return nil
	}
}
