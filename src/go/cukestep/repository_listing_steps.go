package cukestep

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	core "github.com/kkrull/marmot/corerepository"
	support "github.com/kkrull/marmot/cukesupport"
	. "github.com/onsi/gomega"
)

/* Configuration */

// Add step definitions related to listing repositories.
func AddRepositorySteps(ctx *godog.ScenarioContext) {
	ctx.After(func(ctx context.Context, _ *godog.Scenario, err error) (context.Context, error) {
		_thatListing = nil
		return ctx, err
	})

	ctx.Then(`^that repository listing should be empty$`, thatListingShouldBeEmpty)
	ctx.When(`^I list local repositories in that meta repo$`, listLocal)
	ctx.When(`^I list remote repositories in that meta repo$`, listRemote)
}

/* State */

var _thatListing core.Repositories

// Store the most recent listing
func setThatListing(listing core.Repositories) {
	_thatListing = listing
}

// Whichever repositories were returned from the most recent listing, if any.
func thatListing() core.Repositories {
	return _thatListing
}

/* Use */

func listLocal() error {
	if factory, configErr := support.ThatQueryFactory(); configErr != nil {
		return fmt.Errorf("repository_steps: failed to configure; %w", configErr)
	} else if listRepositories, appErr := factory.NewListLocalRepositories(); appErr != nil {
		return fmt.Errorf("repository_steps: failed to initialize; %w", appErr)
	} else if repositories, runErr := listRepositories(); runErr != nil {
		return fmt.Errorf("repository_steps: failed to run query; %w", runErr)
	} else {
		setThatListing(repositories)
		return nil
	}
}

func listRemote() error {
	if factory, configErr := support.ThatQueryFactory(); configErr != nil {
		return fmt.Errorf("repository_steps: failed to configure; %w", configErr)
	} else if listRepositories, appErr := factory.NewListRemoteRepositories(); appErr != nil {
		return fmt.Errorf("repository_steps: failed to initialize; %w", appErr)
	} else if repositories, runErr := listRepositories(); runErr != nil {
		return fmt.Errorf("repository_steps: failed to run query; %w", runErr)
	} else {
		setThatListing(repositories)
		return nil
	}
}

func thatListingShouldBeEmpty() {
	Expect(thatListing().Count()).To(Equal(0))
}

func thatListingShouldHaveLocals(expectedPaths ...string) error {
	return godog.ErrPending
}

func thatListingShouldHaveRemotes(expectedHrefs ...string) {
	remoteUrls := thatListing().RemoteUrls()
	remoteHrefs := make([]string, len(remoteUrls))
	for i, remoteUrl := range remoteUrls {
		remoteHrefs[i] = remoteUrl.String()
	}

	Expect(remoteHrefs).To(ConsistOf(expectedHrefs))
}
