package cukestep

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
)

// State to clear between scenarios
var thoseLocalRepositories []*LocalRepository

func InitLocalRepository(path string) (*LocalRepository, error) {
	repo := &LocalRepository{path: path}
	return repo, repo.Create()
}

type LocalRepository struct {
	path string
}

func (localRepo *LocalRepository) Create() error {
	return os.MkdirAll(localRepo.path, 0o777)
}

func (localRepo *LocalRepository) Delete() error {
	return os.RemoveAll(localRepo.path)
}

// Add step definitions related to repositories on the local filesystem.
func AddLocalRepositorySteps(ctx *godog.ScenarioContext) {
	//TODO KDK: Delete directories containing local repositories, in a Before hook (not After)
	ctx.After(func(ctx context.Context, _ *godog.Scenario, err error) (context.Context, error) {
		var totalErr error = nil
		for _, localRepo := range thoseLocalRepositories {
			totalErr = errors.Join(totalErr, localRepo.Delete())
		}

		thoseLocalRepositories = nil
		return ctx, totalErr
	})

	ctx.Given(`^Git repositories on the local filesystem$`, localGitRepositories)
}

func localGitRepositories() error {
	if testDir, pathErr := support.TestDir(); pathErr != nil {
		return pathErr
	} else if repo, initErr := InitLocalRepository(filepath.Join(testDir, "empty-dir")); initErr != nil {
		return initErr
	} else {
		thoseLocalRepositories = []*LocalRepository{repo}
		return nil
	}
}
