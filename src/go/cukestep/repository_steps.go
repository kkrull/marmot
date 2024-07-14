package cukestep

import (
	"fmt"

	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
	main "github.com/kkrull/marmot/mainfactory"
	. "github.com/onsi/gomega"
)

// Add step definitions related to repositories.
func AddRepositorySteps(ctx *godog.ScenarioContext) {
	ctx.Given(`^I have registered local repositories$`, registerLocalRepositories)
	ctx.When(`^I list repositories in that meta repo$`, listRepositoriesInThatMetaRepo)

	ctx.Then(`^that repository listing should be empty$`, thatRepositoryListingShouldBeEmpty)
	ctx.Then(`^that repository listing should include local repositories that were registered$`, thatRepositoryListingShouldIncludeLocalRepositories)
}

/* List repositories */

var thatRepositoryListing []string //TODO KDK: Does this need to be reset between scenarios too?

func listRepositoriesInThatMetaRepo() error {
	if metaRepoPath, pathErr := support.ThatMetaRepo(); pathErr != nil {
		return fmt.Errorf("repository_steps: %w", pathErr)
	} else if repoList, listErr := listRepositories(metaRepoPath); listErr != nil {
		return fmt.Errorf("repository_steps: failed to list repositories: %w", listErr)
	} else {
		thatRepositoryListing = repoList
		return nil
	}
}

func listRepositories(metaRepoPath string) ([]string, error) {
	factory := &main.CommandFactory{}
	factory.WithJsonFileSource(metaRepoPath)
	if listQuery, factoryErr := factory.ListRepositoriesQuery(); factoryErr != nil {
		return nil, fmt.Errorf("repository_steps: failed to initialize: %w", factoryErr)
	} else if repositories, runErr := listQuery.Run(); runErr != nil {
		return nil, fmt.Errorf("repository_steps: failed to run query: %w", runErr)
	} else {
		return repositories.Names(), nil
	}
}

func thatRepositoryListingShouldBeEmpty() {
	Expect(thatRepositoryListing).To(BeEmpty())
}

func thatRepositoryListingShouldIncludeLocalRepositories() error {
	return godog.ErrPending
}

/* Register repositories */

func registerLocalRepositories() error {
	return godog.ErrPending
}
