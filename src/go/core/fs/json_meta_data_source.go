package fs

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// Stores meta data in JSON files in a directory that Marmot manages
type JsonMetaDataSource struct {
	Path string
}

func (source *JsonMetaDataSource) Init() error {
	_, statErr := os.Stat(source.Path)
	if errors.Is(statErr, fs.ErrNotExist) {
		return createMetaData(source.Path)
	} else if statErr != nil {
		return statErr
	} else {
		return fmt.Errorf("%s: path already exists", source.Path)
	}
}

func createMetaData(path string) error {
	if dirErr := os.MkdirAll(path, fs.ModePerm); dirErr != nil {
		return fmt.Errorf("createMetaData %s: %w", path, dirErr)
	}

	return nil
}
