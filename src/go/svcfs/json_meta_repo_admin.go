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
	return initDirectory(
		localDataFile(repositoryDir),
		metaDataFile(repositoryDir),
		InitMetaRepoData(admin.version))
}

func initDirectory(localDataFile string, metaDataFile string, rootObject *rootObjectData) error {
	metaDataDir := path.Dir(metaDataFile)
	if dirErr := os.MkdirAll(metaDataDir, fs.ModePerm); dirErr != nil {
		return fmt.Errorf("failed to make directory %s; %w", metaDataDir, dirErr)
	} else if writeMetaDataErr := rootObject.WriteTo(metaDataFile); writeMetaDataErr != nil {
		return fmt.Errorf("failed to write shared meta data file %s; %w", metaDataFile, writeMetaDataErr)
	} else {
		return rootObject.WriteTo(localDataFile)
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
