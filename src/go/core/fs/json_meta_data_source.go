package fs

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// Stores meta data in JSON files in a directory that Marmot manages
type JsonMetaDataSource struct {
	Path string
}

func (source *JsonMetaDataSource) Init() error {
	_, statErr := os.Stat(source.Path)
	if errors.Is(statErr, fs.ErrNotExist) {
		return createMetaData(filepath.Join(source.Path, ".marmot"))
	} else if statErr != nil {
		return statErr
	} else {
		return fmt.Errorf("%s: path already exists", source.Path)
	}
}

func createMetaData(metaDataDir string) error {
	if dirErr := os.MkdirAll(metaDataDir, fs.ModePerm); dirErr != nil {
		return fmt.Errorf("createMetaData %s: %w", metaDataDir, dirErr)
	}

	metaDataFile := filepath.Join(metaDataDir, "meta-repo.json")
	if _, fileErr := os.Create(metaDataFile); fileErr != nil {
		return fmt.Errorf("createMetaData %s: %w", metaDataFile, fileErr)
	}

	return nil
}
