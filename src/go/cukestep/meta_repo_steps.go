package cukestep

import (
	"fmt"
	"path/filepath"

	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
	main "github.com/kkrull/marmot/mainfactory"
)

// Add step definitions to manage the life cycle of a meta repo
func AddMetaRepoSteps(ctx *godog.ScenarioContext) {
	ctx.Given(`^I have initialized a new meta repo$`, initializeNewMetaRepo)
}

/* State */

var thatMetaRepo string

func ThatMetaRepo() (string, error) {
	if thatMetaRepo == "" {
		return "", fmt.Errorf("meta_repo_steps: no meta repo has been configured")
	}

	return thatMetaRepo, nil
}

func setThatMetaRepo(path string) {
	thatMetaRepo = path
}

/* Steps */

func initializeNewMetaRepo() error {
	if thatMetaRepo != "" {
		return fmt.Errorf("meta_repo_steps: meta repo has already been configured at %s", thatMetaRepo)
	} else if testDir, mkdirErr := support.TestDir(); mkdirErr != nil {
		return mkdirErr
	} else {
		setThatMetaRepo(filepath.Join(testDir, "meta"))
	}

	cmdFactory := &main.CommandFactory{}
	cmdFactory.WithJsonFileSource(thatMetaRepo)
	if initCmd, factoryErr := cmdFactory.InitCommand(); factoryErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize: %w", factoryErr)
	} else if runErr := initCmd.Run(); runErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize repository %s: %w", thatMetaRepo, runErr)
	} else {
		return nil
	}
}
