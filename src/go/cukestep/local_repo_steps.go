package cukestep

import (
	"context"

	"github.com/cucumber/godog"
	core "github.com/kkrull/marmot/corerepository"
)

// State to clear between scenarios
var thoseLocalRepositories core.Repositories

// Add step definitions related to repositories on the local filesystem.
func AddLocalRepositorySteps(ctx *godog.ScenarioContext) {
	ctx.After(func(ctx context.Context, _ *godog.Scenario, err error) (context.Context, error) {
		//TODO KDK: Delete directories containing local repositories
		thoseLocalRepositories = nil
		return ctx, err
	})

	ctx.Given(`^Git repositories on the local filesystem$`, localGitRepositories)
}

func localGitRepositories() error {
	return godog.ErrPending
}
