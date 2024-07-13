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

func (metaRepo *JsonMetaDataRepo) Init() error {
	_, statErr := os.Stat(metaRepo.repositoryDir)
	if errors.Is(statErr, fs.ErrNotExist) {
		return metaRepo.createMetaData()
	} else if statErr != nil {
		return statErr
	} else {
		return fmt.Errorf("%s: path already exists", metaRepo.repositoryDir)
	}
}

func (metaRepo *JsonMetaDataRepo) createMetaData() error {
	if dirErr := os.MkdirAll(metaRepo.metaDataDir, fs.ModePerm); dirErr != nil {
		return fmt.Errorf("createMetaData %s: %w", metaRepo.metaDataDir, dirErr)
	} else if _, fileErr := os.Create(metaRepo.metaDataFile); fileErr != nil {
		return fmt.Errorf("createMetaData %s: %w", metaRepo.metaDataFile, fileErr)
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
