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
	stat, statErr := os.Stat(source.Path)
	if errors.Is(statErr, fs.ErrNotExist) {
		fmt.Printf("[prod] creating directories: %s\n", source.Path)
		dirErr := os.MkdirAll(source.Path, fs.ModePerm)
		if dirErr != nil {
			fmt.Printf("[prod] failed to make directories: %s\n", source.Path)
			fmt.Println(dirErr.Error())
		}

		return nil
	} else if statErr != nil {
		return statErr
	} else if stat != nil {
		return fmt.Errorf("%s: path already exists", source.Path)
	}

	return fmt.Errorf("%s: undefined behavior", source.Path)
}
