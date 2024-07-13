package step_definitions

import (
	"fmt"

	"github.com/cucumber/godog"
	. "github.com/onsi/gomega"
)

// Add step definitions related to repositories
func AddRepositorySteps(ctx *godog.ScenarioContext) {
	ctx.When(`^I list repositories in that meta repo$`, listRepositoriesInThatMetaRepo)
	ctx.Then(`^that repository listing should be empty$`, thatRepositoryListingShouldBeEmpty)
}

/* Listing repositories */

var thatRepositoryListing []string

func listRepositoriesInThatMetaRepo() error {
	if metaRepoPath, metaRepoPathErr := ThatMetaRepo(); metaRepoPathErr != nil {
		return metaRepoPathErr
	} else if repoList, repoListErr := listRepositories(metaRepoPath); repoListErr != nil {
		return repoListErr
	} else {
		thatRepositoryListing = repoList
		return nil
	}
}

func listRepositories(metaRepoPath string) ([]string, error) {
	fmt.Printf("[repository_steps] listing repositories in %s\n", metaRepoPath)
	return nil, godog.ErrPending
}

func thatRepositoryListingShouldBeEmpty() {
	Expect(thatRepositoryListing).To(BeEmpty())
}
