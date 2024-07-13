package cukestep

import (
	"fmt"

	"github.com/cucumber/godog"
	"github.com/kkrull/marmot/cukesupport"
	main "github.com/kkrull/marmot/mainfactory"
	. "github.com/onsi/gomega"
)

// Add step definitions related to repositories.
func AddRepositorySteps(ctx *godog.ScenarioContext) {
	ctx.When(`^I list repositories in that meta repo$`, listRepositoriesInThatMetaRepo)
	ctx.Then(`^that repository listing should be empty$`, thatRepositoryListingShouldBeEmpty)
}

/* Listing repositories */

var thatRepositoryListing []string

func listRepositoriesInThatMetaRepo() error {
	if metaRepoPath, metaRepoPathErr := cukesupport.ThatMetaRepo(); metaRepoPathErr != nil {
		return metaRepoPathErr
	} else if repoList, repoListErr := listRepositories(metaRepoPath); repoListErr != nil {
		return repoListErr
	} else {
		thatRepositoryListing = repoList
		return nil
	}
}

func listRepositories(metaRepoPath string) ([]string, error) {
	factory := &main.CommandFactory{}
	if thatMetaRepo, configErr := cukesupport.ThatMetaRepo(); configErr != nil {
		return nil, fmt.Errorf("repository_steps: failed to configure: %w", configErr)
	} else {
		factory.WithJsonFileSource(thatMetaRepo)
	}

	if listQuery, factoryErr := factory.ListRepositoriesQuery(); factoryErr != nil {
		return nil, fmt.Errorf("repository_steps: failed to initialize: %w", factoryErr)
	} else if repositories, listErr := listQuery.Run(); listErr != nil {
		return nil, fmt.Errorf("repository_steps: failed to run query: %w", listErr)
	} else {
		return repositories.Names(), nil
	}
}

func thatRepositoryListingShouldBeEmpty() {
	Expect(thatRepositoryListing).To(BeEmpty())
}
