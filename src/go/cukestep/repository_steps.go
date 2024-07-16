package cukestep

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
	main "github.com/kkrull/marmot/mainfactory"
	. "github.com/onsi/gomega"
)

// State to clear between scenarios
var thatListing []string

// Add step definitions related to repositories.
func AddRepositorySteps(ctx *godog.ScenarioContext) {
	ctx.After(func(ctx context.Context, _ *godog.Scenario, err error) (context.Context, error) {
		thatListing = nil
		return ctx, err
	})

	ctx.Given(`^I have registered remote repositories$`, registerRemote)
	ctx.When(`^I list repositories in that meta repo$`, listInThatMetaRepo)

	ctx.Then(`^that repository listing should be empty$`, thatListingShouldBeEmpty)
	ctx.Then(`^that repository listing should include those remote repositories$`, thatListingShouldHaveRemotes)
}

/* List repositories */

func listInThatMetaRepo() error {
	if metaRepoPath, pathErr := support.ThatMetaRepo(); pathErr != nil {
		return fmt.Errorf("repository_steps: failed to configure; %w", pathErr)
	} else if repoList, listErr := listRepositories(metaRepoPath); listErr != nil {
		return fmt.Errorf("repository_steps: failed to list repositories; %w", listErr)
	} else {
		thatListing = repoList
		return nil
	}
}

func listRepositories(metaRepoPath string) ([]string, error) {
	factory := &main.CommandFactory{}
	factory.WithJsonFileSource(metaRepoPath)
	if listQuery, factoryErr := factory.ListRepositoriesQuery(); factoryErr != nil {
		return nil, fmt.Errorf("repository_steps: failed to initialize; %w", factoryErr)
	} else if repositories, runErr := listQuery.Run(); runErr != nil {
		return nil, fmt.Errorf("repository_steps: failed to run query; %w", runErr)
	} else {
		return repositories.Names(), nil
	}
}

func thatListingShouldBeEmpty() {
	Expect(thatListing).To(BeEmpty())
}

func thatListingShouldHaveRemotes() error {
	return godog.ErrPending
}

/* Register repositories */

func registerRemote() error {
	return godog.ErrPending
}
