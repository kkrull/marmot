package cukestep

import (
	"fmt"

	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
)

// Add step definitions to manage the life cycle of a meta repo.
func AddMetaRepoSteps(ctx *godog.ScenarioContext) {
	initNewMetaRepoSC := func() error { return initNewMetaRepo(ctx) }
	ctx.Given(`^I have initialized a new meta repo$`, initNewMetaRepoSC)
}

/* Steps */

func initNewMetaRepo(ctx *godog.ScenarioContext) error {
	factory := support.ThatCommandFactoryS("42")
	if initCmd, factoryErr := factory.NewInitMetaRepo(); factoryErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize; %w", factoryErr)
	} else if thatMetaRepo, initErr := support.InitThatMetaRepo(ctx); initErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize path to meta repo; %w", initErr)
	} else if runErr := initCmd.Run(thatMetaRepo); runErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize repository; %w", runErr)
	} else {
		return nil
	}
}
