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
	}
	// if err != nil {
	// 	return err
	// }

	if stat != nil && stat.IsDir() {
		return errors.New(fmt.Sprintf("%s: path already exists", source.Path))
	}

	return nil
}
