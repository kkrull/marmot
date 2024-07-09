package fs

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func NewJsonMetaDataSource(repositoryPath string) *JsonMetaDataSource {
	return &JsonMetaDataSource{RepositoryPath: repositoryPath}
}

// Stores meta data in JSON files in a directory that Marmot manages
type JsonMetaDataSource struct {
	RepositoryPath string
}

func (source *JsonMetaDataSource) Init() error {
	_, statErr := os.Stat(source.RepositoryPath)
	if errors.Is(statErr, fs.ErrNotExist) {
		return createMetaData(filepath.Join(source.RepositoryPath, ".marmot"))
	} else if statErr != nil {
		return statErr
	} else {
		return fmt.Errorf("%s: path already exists", source.RepositoryPath)
	}
}

func createMetaData(metaDataDir string) error {
	metaDataFile := filepath.Join(metaDataDir, "meta-repo.json")
	if dirErr := os.MkdirAll(metaDataDir, fs.ModePerm); dirErr != nil {
		return fmt.Errorf("createMetaData %s: %w", metaDataDir, dirErr)
	} else if _, fileErr := os.Create(metaDataFile); fileErr != nil {
		return fmt.Errorf("createMetaData %s: %w", metaDataFile, fileErr)
	}

	return nil
}
