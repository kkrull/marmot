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
		//Ignore an existing meta repo, for now
		return nil
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

func (admin *JsonMetaRepoAdmin) IsMetaRepo(path string) (bool, error) {
	pathStat, pathErr := os.Stat(path)
	if errors.Is(pathErr, fs.ErrNotExist) {
		return false, nil
	} else if pathErr != nil {
		return false, fmt.Errorf("%s: failed to stat meta repo path; %w", path, pathErr)
	} else if pathStat.Mode().IsRegular() {
		return false, nil
	}

	marmotDir := metaDataDir(path)
	_, marmotDirErr := os.Stat(marmotDir)
	if errors.Is(marmotDirErr, fs.ErrNotExist) {
		return false, nil
	} else if marmotDirErr != nil {
		return false, fmt.Errorf("%s: failed to stat Marmot directory; %w", marmotDir, marmotDirErr)
	} else {
		return true, nil
	}
}
