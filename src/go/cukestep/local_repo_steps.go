package cukestep

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
)

// State to clear between scenarios
var thoseLocalRepositories *support.LocalRepositories = support.NoLocalRepositories()

// Add step definitions related to repositories on the local filesystem.
func AddLocalRepositorySteps(ctx *godog.ScenarioContext) {
	ctx.After(func(ctx context.Context, _ *godog.Scenario, err error) (context.Context, error) {
		if err != nil {
			return ctx, err
		} else {
			return ctx, thoseLocalRepositories.DeleteAll()
		}
	})

	ctx.Given(`^Git repositories on the local filesystem$`, createLocalGitRepositories)
	ctx.Given(`^I have registered those local repositories with a meta repo$`, registerLocal)

	ctx.When(`^I list local repositories in that meta repo$`, listLocal)
	ctx.Then(`^that repository listing should include those local repositories$`, thatListingShouldHaveLocals)
}

func createLocalGitRepositories() error {
	if repoDir, pathErr := support.TestSubDir("empty-dir"); pathErr != nil {
		return pathErr
	} else if repo, repoErr := support.InitLocalRepository(repoDir); repoErr != nil {
		return repoErr
	} else {
		thoseLocalRepositories = support.SomeLocalRepositories(repo)
		return nil
	}
}

func listLocal() error {
	if factory, configErr := support.ThatQueryFactory(); configErr != nil {
		return fmt.Errorf("repository_steps: failed to configure; %w", configErr)
	} else if listRepositories, appErr := factory.NewListLocalRepositories(); appErr != nil {
		return fmt.Errorf("repository_steps: failed to initialize; %w", appErr)
	} else if repositories, runErr := listRepositories(); runErr != nil {
		return fmt.Errorf("repository_steps: failed to run query; %w", runErr)
	} else {
		thatListing = repositories
		return nil
	}
}

func registerLocal() error {
	return godog.ErrPending
}

// TODO KDK: Just implement the application command to register and the application query; leave CLI for another PR
func thatListingShouldHaveLocals() error {
	return godog.ErrPending
}
