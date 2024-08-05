package cukesupport

import (
	"errors"
	"os"
	"path/filepath"
)

// Initialize a Git repository at the specified path on the local filesystem.
func InitLocalRepository(path string) (*LocalRepository, error) {
	if absPath, pathErr := filepath.Abs(path); pathErr != nil {
		return nil, pathErr
	} else {
		repo := &LocalRepository{path: absPath}
		return repo, repo.Create()
	}
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

/* Container */

// Create an empty container for local repositories.
func NoLocalRepositories() *LocalRepositories {
	return &LocalRepositories{repositories: make([]*LocalRepository, 0)}
}

// Create a container for local repositories.
func SomeLocalRepositories(repositories ...*LocalRepository) *LocalRepositories {
	return &LocalRepositories{repositories: repositories}
}

type LocalRepositories struct {
	repositories []*LocalRepository
}

// Exhaustively try to delete each local repository, joining any errors from individual attempts.
func (container *LocalRepositories) DeleteAll() error {
	var totalErr error = nil
	for _, localRepo := range container.repositories {
		totalErr = errors.Join(totalErr, localRepo.Delete())
	}

	return totalErr
}

// Paths on the local filesystem, to each repository.
func (container *LocalRepositories) LocalPaths() []string {
	localPaths := make([]string, len(container.repositories))
	for i, repository := range container.repositories {
		localPaths[i] = repository.path
	}

	return localPaths
}
