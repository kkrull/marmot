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
	emptyFile := EmptyMetaRepoFile("0.0.1")
	if dirErr := os.MkdirAll(meta.metaDataDir, fs.ModePerm); dirErr != nil {
		return fmt.Errorf("failed to make directory %s; %w", meta.metaDataDir, dirErr)
	} else if writeErr := emptyFile.WriteTo(meta.metaDataFile); writeErr != nil {
		return fmt.Errorf("failed to write file %s; %w", meta.metaDataFile, writeErr)
	} else {
		return nil
	}
}

/* RepositorySource */

func (metaRepo *JsonMetaDataRepo) List() (corerepository.Repositories, error) {
	if content, readErr := ReadMetaRepoFile(metaRepo.metaDataFile); readErr != nil {
		return nil, fmt.Errorf("failed to read file %s; %w", metaRepo.metaDataFile, readErr)
	} else if repositories, convertErr := content.ToCoreRepositories(); convertErr != nil {
		return nil, fmt.Errorf("failed to map to core model; %w", convertErr)
	} else {
		return repositories, nil
	}
}

func (metaRepo *JsonMetaDataRepo) RegisterRemote(hostUrl *url.URL) error {
	var encoder *json.Encoder
	if metaDataFd, openErr := os.OpenFile(metaRepo.metaDataFile, os.O_WRONLY, os.ModePerm); openErr != nil {
		return fmt.Errorf("failed to open file %s; %w", metaRepo.metaDataFile, openErr)
	} else {
		encoder = json.NewEncoder(metaDataFd)
	}

	// TODO KDK: Read or remember other registered repositories, instead of just writing the new one
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
