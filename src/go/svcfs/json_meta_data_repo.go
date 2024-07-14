package svcfs

import (
	"errors"
	"fmt"
	"io/fs"
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
		return meta.createMetaData()
	} else if statErr != nil {
		return fmt.Errorf("failed to check for existing meta repo %s; %w", meta.repositoryDir, statErr)
	} else {
		return fmt.Errorf("path already exists: %s", meta.repositoryDir)
	}
}

func (meta *JsonMetaDataRepo) createMetaData() error {
	if dirErr := os.MkdirAll(meta.metaDataDir, fs.ModePerm); dirErr != nil {
		return fmt.Errorf("failed to make directory %s; %w", meta.metaDataDir, dirErr)
	} else if _, fileErr := os.Create(meta.metaDataFile); fileErr != nil {
		return fmt.Errorf("failed to create file %s; %w", meta.metaDataFile, fileErr)
	}

	return nil
}

/* RepositorySource */

func (metaRepo *JsonMetaDataRepo) List() (corerepository.Repositories, error) {
	repositories := &corerepository.RepositoriesArray{
		Repositories: make([]corerepository.Repository, 0),
	}

	return repositories, nil
}
