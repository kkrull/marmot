package svcfs

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	core "github.com/kkrull/marmot/coremetarepo"
)

// A meta repo that stores meta data in JSON files on the file system.
func JsonMetaDataRepo(repositoryPath string) core.MetaDataAdmin {
	metaDataDir := filepath.Join(repositoryPath, ".marmot")
	return &jsonMetaDataRepo{
		repositoryDir: repositoryPath,
		metaDataDir:   metaDataDir,
		metaDataFile:  filepath.Join(metaDataDir, "meta-repo.json"),
	}
}

type jsonMetaDataRepo struct {
	metaDataDir   string
	metaDataFile  string
	repositoryDir string
}

func (source *jsonMetaDataRepo) Init() error {
	_, statErr := os.Stat(source.repositoryDir)
	if errors.Is(statErr, fs.ErrNotExist) {
		return source.createMetaData()
	} else if statErr != nil {
		return statErr
	} else {
		return fmt.Errorf("%s: path already exists", source.repositoryDir)
	}
}

func (source *jsonMetaDataRepo) createMetaData() error {
	if dirErr := os.MkdirAll(source.metaDataDir, fs.ModePerm); dirErr != nil {
		return fmt.Errorf("createMetaData %s: %w", source.metaDataDir, dirErr)
	} else if _, fileErr := os.Create(source.metaDataFile); fileErr != nil {
		return fmt.Errorf("createMetaData %s: %w", source.metaDataFile, fileErr)
	}

	return nil
}
