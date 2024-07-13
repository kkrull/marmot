package cukestep

import (
	"fmt"
	"path/filepath"

	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
	main "github.com/kkrull/marmot/mainfactory"
)

// Add step definitions to manage the life cycle of a meta repo.
func AddMetaRepoSteps(ctx *godog.ScenarioContext) {
	ctx.Given(`^I have initialized a new meta repo$`, initializeNewMetaRepo)
}

/* Steps */

func initializeNewMetaRepo() error {
	var thatMetaRepo string
	if existingMetaRepo := support.PeekThatMetaRepo(); existingMetaRepo != "" {
		return fmt.Errorf("meta_repo_steps: meta repo has already been configured at %s", existingMetaRepo)
	} else if testDir, mkdirErr := support.TestDir(); mkdirErr != nil {
		return mkdirErr
	} else {
		thatMetaRepo = support.SetThatMetaRepo(filepath.Join(testDir, "meta"))
	}

	factory := &main.CommandFactory{}
	factory.WithJsonFileSource(thatMetaRepo)
	if initCmd, factoryErr := factory.InitCommand(); factoryErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize: %w", factoryErr)
	} else if runErr := initCmd.Run(); runErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize repository: %w", runErr)
	} else {
		return nil
	}
}
