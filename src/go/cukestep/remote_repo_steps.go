package cukestep

import (
	"fmt"
	"net/url"

	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
	. "github.com/onsi/gomega"
)

// Add step definitions related to remote repositories.
func AddRemoteRepositorySteps(ctx *godog.ScenarioContext) {
	ctx.Given(`^I have registered remote repositories$`, registerRemote)
	ctx.When(`^I list remote repositories in that meta repo$`, listRemote)
	ctx.Then(`^that repository listing should include those remote repositories$`, thatListingShouldHaveRemotes)
}

func listRemote() error {
	if factory, configErr := support.ThatQueryFactory(); configErr != nil {
		return fmt.Errorf("repository_steps: failed to configure; %w", configErr)
	} else if listRepositories, appErr := factory.NewListRemoteRepositories(); appErr != nil {
		return fmt.Errorf("repository_steps: failed to initialize; %w", appErr)
	} else if repositories, runErr := listRepositories(); runErr != nil {
		return fmt.Errorf("repository_steps: failed to run query; %w", runErr)
	} else {
		thatListing = repositories
		return nil
	}
}

func registerRemote() error {
	if remoteUrl, parseErr := url.Parse("https://github.com/actions/checkout"); parseErr != nil {
		return parseErr
	} else if factory, factoryErr := support.ThatCommandFactory(); factoryErr != nil {
		return fmt.Errorf("repository_steps: failed to configure; %w", factoryErr)
	} else if registerCmd, factoryErr := factory.NewRegisterRemoteRepositories(); factoryErr != nil {
		return fmt.Errorf("repository_steps: failed to initialize; %w", factoryErr)
	} else if runErr := registerCmd.Run([]*url.URL{remoteUrl}); runErr != nil {
		return fmt.Errorf("repository_steps: failed to register repositories; %w", runErr)
	} else {
		return nil
	}
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
