package svcfs

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func NewJsonMetaRepoAdmin() *JsonMetaRepoAdmin {
	return &JsonMetaRepoAdmin{}
}

// Creates meta repos that store meta data in JSON files on the local file system.
type JsonMetaRepoAdmin struct{}

/* MetaDataAdmin */

func (*JsonMetaRepoAdmin) Create(repositoryDir string) error {
	_, statErr := os.Stat(repositoryDir)
	if errors.Is(statErr, fs.ErrNotExist) {
		return initDirectory(repositoryDir)
	} else if statErr != nil {
		return fmt.Errorf("failed to check for existing meta repo %s; %w", repositoryDir, statErr)
	} else {
		return fmt.Errorf("path already exists: %s", repositoryDir)
	}
}

func initDirectory(repositoryDir string) error {
	metaDataDir := filepath.Join(repositoryDir, ".marmot")
	metaDataFile := filepath.Join(metaDataDir, "meta-repo.json")

	emptyFile := EmptyMetaRepoFile("0.0.1")
	if dirErr := os.MkdirAll(metaDataDir, fs.ModePerm); dirErr != nil {
		return fmt.Errorf("failed to make directory %s; %w", metaDataDir, dirErr)
	} else if writeErr := emptyFile.WriteTo(metaDataFile); writeErr != nil {
		return fmt.Errorf("failed to write file %s; %w", metaDataFile, writeErr)
	} else {
		return nil
	}
}
