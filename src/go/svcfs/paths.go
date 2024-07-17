package svcfs

import "path/filepath"

// Path to the general directory where Marmot stores its data.
func metaDataDir(metaRepoDir string) string {
	return filepath.Join(metaRepoDir, ".marmot")
}

// Path to the specific file where Marmot store its data.
func metaDataFile(metaRepoDir string) string {
	return filepath.Join(metaDataDir(metaRepoDir), "meta-repo.json")
}
