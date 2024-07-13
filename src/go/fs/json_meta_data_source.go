package fs

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	metarepo "github.com/kkrull/marmot/core-metarepo"
)

// Makes a JsonMetaDataSource for the meta repository at repositoryPath, using conventional paths to
// meta data within it.
func NewJsonMetaDataSource(repositoryPath string) metarepo.MetaDataSource {
	metaDataDir := filepath.Join(repositoryPath, ".marmot")
	return &JsonMetaDataSource{
		repositoryDir: repositoryPath,
		metaDataDir:   metaDataDir,
		metaDataFile:  filepath.Join(metaDataDir, "meta-repo.json"),
	}
}

// Stores meta data in JSON files in a directory that Marmot manages
type JsonMetaDataSource struct {
	metaDataDir   string
	metaDataFile  string
	repositoryDir string
}

func (source *JsonMetaDataSource) Init() error {
	_, statErr := os.Stat(source.repositoryDir)
	if errors.Is(statErr, fs.ErrNotExist) {
		return source.createMetaData()
	} else if statErr != nil {
		return statErr
	} else {
		return fmt.Errorf("%s: path already exists", source.repositoryDir)
	}
}

func (source *JsonMetaDataSource) createMetaData() error {
	if dirErr := os.MkdirAll(source.metaDataDir, fs.ModePerm); dirErr != nil {
		return fmt.Errorf("createMetaData %s: %w", source.metaDataDir, dirErr)
	} else if _, fileErr := os.Create(source.metaDataFile); fileErr != nil {
		return fmt.Errorf("createMetaData %s: %w", source.metaDataFile, fileErr)
	}

	return nil
}
