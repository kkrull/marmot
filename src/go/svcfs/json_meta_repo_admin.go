package svcfs

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
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
		return initDirectory(metaDataDir(repositoryDir), metaDataFile(repositoryDir))
	} else if statErr != nil {
		return fmt.Errorf("failed to check for existing meta repo %s; %w", repositoryDir, statErr)
	} else {
		return fmt.Errorf("path already exists: %s", repositoryDir)
	}
}

func initDirectory(metaDataDirS string, metaDataFileS string) error {
	emptyFile := EmptyMetaRepoFile("0.0.1") //TODO KDK: Get the real version number
	if dirErr := os.MkdirAll(metaDataDirS, fs.ModePerm); dirErr != nil {
		return fmt.Errorf("failed to make directory %s; %w", metaDataDirS, dirErr)
	} else if writeErr := emptyFile.WriteTo(metaDataFileS); writeErr != nil {
		return fmt.Errorf("failed to write file %s; %w", metaDataFileS, writeErr)
	} else {
		return nil
	}
}
