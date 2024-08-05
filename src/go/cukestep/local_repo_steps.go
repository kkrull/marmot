package cukestep

import (
	"context"
	"errors"

	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
)

// State to clear between scenarios
var thoseLocalRepositories []*support.LocalRepository

// Add step definitions related to repositories on the local filesystem.
func AddLocalRepositorySteps(ctx *godog.ScenarioContext) {
	ctx.After(func(ctx context.Context, _ *godog.Scenario, err error) (context.Context, error) {
		var totalErr error = err
		for _, localRepo := range thoseLocalRepositories {
			totalErr = errors.Join(totalErr, localRepo.Delete())
		}

		thoseLocalRepositories = nil
		return ctx, totalErr
	})

	ctx.Given(`^Git repositories on the local filesystem$`, localGitRepositories)
}

func localGitRepositories() error {
	if repoDir, pathErr := support.TestSubDir("empty-dir"); pathErr != nil {
		return pathErr
	} else if repo, repoErr := support.InitLocalRepository(repoDir); repoErr != nil {
		return repoErr
	} else {
		thoseLocalRepositories = []*support.LocalRepository{repo}
		return nil
	}
}
