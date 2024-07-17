package cukestep

import (
	"fmt"

	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
	main "github.com/kkrull/marmot/mainfactory"
)

// Add step definitions to manage the life cycle of a meta repo.
func AddMetaRepoSteps(ctx *godog.ScenarioContext) {
	initNewMetaRepoSC := func() error { return initNewMetaRepo(ctx) }
	ctx.Given(`^I have initialized a new meta repo$`, initNewMetaRepoSC)
}

/* Steps */

func initNewMetaRepo(ctx *godog.ScenarioContext) error {
	thatMetaRepo, initErr := support.InitThatMetaRepo(ctx)
	if initErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize path to meta repo; %w", initErr)
	}

	factory := &main.CommandFactory{}
	factory.WithJsonFileSource(thatMetaRepo)
	if initCmd, factoryErr := factory.InitCommand(); factoryErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize; %w", factoryErr)
	} else if runErr := initCmd.RunP(thatMetaRepo); runErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize repository; %w", runErr)
	} else {
		return nil
	}
}
