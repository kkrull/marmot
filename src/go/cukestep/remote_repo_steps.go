package cukestep

import (
	"fmt"
	"net/url"

	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
)

/* Configuration */

// Add step definitions related to remote repositories.
func AddRemoteRepositorySteps(ctx *godog.ScenarioContext) {
	ctx.Given(`^I have registered remote repositories with a meta repo$`, func() error {
		return registerRemote("https://github.com/actions/checkout")
	})

	ctx.Then(`^that repository listing should include those remote repositories$`, func() {
		thatListingShouldHaveRemotes("https://github.com/actions/checkout")
	})
}

/* Steps */

func registerRemote(remoteHref string) error {
	if remoteUrl, parseErr := url.Parse(remoteHref); parseErr != nil {
		return parseErr
	} else if factory, factoryErr := support.ThatCommandFactory(); factoryErr != nil {
		return fmt.Errorf("remote_repo_steps: failed to configure; %w", factoryErr)
	} else if registerCmd, factoryErr := factory.NewRegisterRemoteRepositories(); factoryErr != nil {
		return fmt.Errorf("remote_repo_steps: failed to initialize; %w", factoryErr)
	} else {
		return registerCmd.Run([]*url.URL{remoteUrl})
	}
}
