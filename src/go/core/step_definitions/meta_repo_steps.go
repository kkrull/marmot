package step_definitions

import (
	"fmt"
	"path/filepath"

	"github.com/cucumber/godog"
	"github.com/kkrull/marmot-core/cmd"
	"github.com/kkrull/marmot-core/fs"
	"github.com/kkrull/marmot-core/hooks"
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
	} else if testDir, mkdirErr := hooks.TestDir(); mkdirErr != nil {
		return mkdirErr
	} else {
		setThatMetaRepo(filepath.Join(testDir, "meta"))
	}

	jsonDataSource := fs.NewJsonMetaDataSource(thatMetaRepo)
	initCmd := &cmd.InitCommand{MetaDataSource: jsonDataSource}
	if runErr := initCmd.Run(); runErr != nil {
		return fmt.Errorf("failed to initialize repository %s: %w", thatMetaRepo, runErr)
	} else {
		return nil
	}
}
