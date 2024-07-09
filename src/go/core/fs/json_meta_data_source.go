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
	stat, err := os.Stat(source.Path)
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	} else if err != nil {
		return err
	} else if stat != nil && stat.IsDir() {
		return fmt.Errorf("%s: path already exists", source.Path)
	}

	return nil
}
