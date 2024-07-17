package cukestep

import (
	"context"
	"fmt"
	"net/url"

	"github.com/cucumber/godog"
	core "github.com/kkrull/marmot/corerepository"
	support "github.com/kkrull/marmot/cukesupport"
	main "github.com/kkrull/marmot/mainfactory"
	. "github.com/onsi/gomega"
)

// State to clear between scenarios
var thatListing core.Repositories

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
	if factory, factoryErr := factoryForThatMetaRepo(); factoryErr != nil {
		return fmt.Errorf("repository_steps: failed to configure; %w", factoryErr)
	} else if listRepositories, factoryErr := factory.ListRemoteRepositoriesQuery(); factoryErr != nil {
		return fmt.Errorf("repository_steps: failed to initialize; %w", factoryErr)
	} else if repositories, runErr := listRepositories(); runErr != nil {
		return fmt.Errorf("repository_steps: failed to run query; %w", runErr)
	} else {
		thatListing = repositories
		return nil
	}
}

func thatListingShouldBeEmpty() {
	Expect(thatListing.Count()).To(Equal(0))
}

func thatListingShouldHaveRemotes() error {
	remoteUrls := thatListing.RemoteUrls()
	remoteHrefs := make([]string, len(remoteUrls))
	for i, remoteUrl := range remoteUrls {
		remoteHrefs[i] = remoteUrl.String()
	}

	Expect(remoteHrefs).To(ConsistOf("https://github.com/actions/checkout"))
	return nil
}

/* Register repositories */

func registerRemote() error {
	if remoteUrl, parseErr := url.Parse("https://github.com/actions/checkout"); parseErr != nil {
		return parseErr
	} else if factory, factoryErr := factoryForThatMetaRepo(); factoryErr != nil {
		return fmt.Errorf("repository_steps: failed to configure; %w", factoryErr)
	} else if registerCmd, factoryErr := factory.RegisterRemoteRepositoriesCommand(); factoryErr != nil {
		return fmt.Errorf("repository_steps: failed to initialize; %w", factoryErr)
	} else if runErr := registerCmd.Run([]*url.URL{remoteUrl}); runErr != nil {
		return fmt.Errorf("repository_steps: failed to register repositories; %w", runErr)
	} else {
		return nil
	}
}

/* Configuration */

func factoryForThatMetaRepo() (*main.CommandQueryFactory, error) {
	if metaRepoPath, pathErr := support.ThatMetaRepo(); pathErr != nil {
		return nil, fmt.Errorf("repository_steps: failed to configure; %w", pathErr)
	} else {
		factory := &main.CommandQueryFactory{}
		factory.ForLocalMetaRepo(metaRepoPath)
		return factory, nil
	}
}
