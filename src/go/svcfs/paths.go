package svcfs

import "path/filepath"

// Path to the general directory where Marmot stores its data.
func metaDataDir(metaRepoDir string) string {
	return filepath.Join(metaRepoDir, ".marmot")
}

// Path to the specific file where Marmot stores machine-specific data.
func localDataFile(metaRepoDir string) string {
	return filepath.Join(metaDataDir(metaRepoDir), "meta-repo-local.json")
}

// Path to the specific file where Marmot stores shared data.
func sharedDataFile(metaRepoDir string) string {
	return filepath.Join(metaDataDir(metaRepoDir), "meta-repo.json")
}
