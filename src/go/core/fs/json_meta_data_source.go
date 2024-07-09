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
	dirErr := os.MkdirAll(path, fs.ModePerm)
	if dirErr != nil {
		fmt.Printf("[prod] failed to make directories: %s\n", path)
		fmt.Println(dirErr.Error())
	}

	return nil
}
