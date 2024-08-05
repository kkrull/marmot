package cukestep

import (
	"fmt"
	"net/url"

	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
	. "github.com/onsi/gomega"
)

/* Configuration */

// Add step definitions related to remote repositories.
func AddRemoteRepositorySteps(ctx *godog.ScenarioContext) {
	ctx.Given(`^I have registered remote repositories$`, registerRemote)
	ctx.When(`^I list remote repositories in that meta repo$`, listRemote)
	ctx.Then(`^that repository listing should include those remote repositories$`, thatListingShouldHaveRemotes)
}

/* Steps */

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
	remoteUrls := thatListing().RemoteUrls()
	remoteHrefs := make([]string, len(remoteUrls))
	for i, remoteUrl := range remoteUrls {
		remoteHrefs[i] = remoteUrl.String()
	}

	Expect(remoteHrefs).To(ConsistOf("https://github.com/actions/checkout"))
	return nil
}
