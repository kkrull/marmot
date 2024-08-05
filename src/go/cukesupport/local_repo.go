package cukesupport

import "os"

// Initialize a Git repository at the specified path on the local filesystem.
func InitLocalRepository(path string) (*LocalRepository, error) {
	repo := &LocalRepository{path: path}
	return repo, repo.Create()
}

// A Git repository on the local filesystem.
type LocalRepository struct {
	path string
}

// Create the Git repository.
func (localRepo *LocalRepository) Create() error {
	return os.MkdirAll(localRepo.path, 0o777)
}

// Delete the directory containing the Git repository.
func (localRepo *LocalRepository) Delete() error {
	return os.RemoveAll(localRepo.path)
}
