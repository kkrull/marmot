package svcfs

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
)

func NewJsonMetaRepoAdmin(version string) *JsonMetaRepoAdmin {
	return &JsonMetaRepoAdmin{version: version}
}

// Creates meta repos that store meta data in JSON files on the local file system.
type JsonMetaRepoAdmin struct {
	version string
}

/* MetaDataAdmin */

func (admin *JsonMetaRepoAdmin) Create(repositoryDir string) error {
	marmotDataDir := metaDataDir(repositoryDir)
	_, statErr := os.Stat(marmotDataDir)
	if errors.Is(statErr, fs.ErrNotExist) {
		return initDirectory(metaDataFile(repositoryDir), InitMetaRepoData(admin.version))
	} else if statErr != nil {
		return fmt.Errorf("failed to check for existing meta repo %s; %w", repositoryDir, statErr)
	} else {
		return fmt.Errorf("path already exists: %s", marmotDataDir)
	}
}

func initDirectory(metaDataFile string, rootObject *rootObjectData) error {
	metaDataDir := path.Dir(metaDataFile)
	if dirErr := os.MkdirAll(metaDataDir, fs.ModePerm); dirErr != nil {
		return fmt.Errorf("failed to make directory %s; %w", metaDataDir, dirErr)
	} else if writeErr := rootObject.WriteTo(metaDataFile); writeErr != nil {
		return fmt.Errorf("failed to write file %s; %w", metaDataFile, writeErr)
	} else {
		return nil
	}
}
