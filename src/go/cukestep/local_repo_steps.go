package cukestep

import (
	"context"

	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
)

// State to clear between scenarios
var thoseLocalRepositories *support.LocalRepositories = support.NoLocalRepositories()

// Add step definitions related to repositories on the local filesystem.
func AddLocalRepositorySteps(ctx *godog.ScenarioContext) {
	ctx.Given(`^Git repositories on the local filesystem$`, localGitRepositories)
	ctx.After(func(ctx context.Context, _ *godog.Scenario, err error) (context.Context, error) {
		if err != nil {
			return ctx, err
		} else {
			return ctx, thoseLocalRepositories.DeleteAll()
		}
	})
}

func localGitRepositories() error {
	if repoDir, pathErr := support.TestSubDir("empty-dir"); pathErr != nil {
		return pathErr
	} else if repo, repoErr := support.InitLocalRepository(repoDir); repoErr != nil {
		return repoErr
	} else {
		thoseLocalRepositories = support.SomeLocalRepositories(repo)
		return nil
	}
}
