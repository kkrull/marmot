package cukesupport

import (
	"errors"
	"os"
)

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

/* Container */

func NoLocalRepositories() *LocalRepositories {
	return &LocalRepositories{repositories: make([]*LocalRepository, 0)}
}

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
