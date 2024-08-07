package cukesupport

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/cucumber/godog"
)

// The name of the tag to identify in feature files
const MetaRepoTag = "@MetaRepo"

// State to clear between scenarios
var _thatMetaRepo string

// Add fixtures to reset state between scenarios.
func addMetaRepoFixture(ctx *godog.ScenarioContext) {
	ctx.After(afterMetaRepo)
	ctx.Before(beforeMetaRepo)
}

func afterMetaRepo(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
	forgetThatMetaRepo()
	return ctx, err
}

func beforeMetaRepo(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	return ctx, nil
}

func forgetThatMetaRepo() {
	_thatMetaRepo = ""
}

// Convenience method initializing a meta repo with this version.
func InitNewMetaRepoS(version string) error {
	factory := ThatCommandFactoryS(version)
	if initCmd, factoryErr := factory.NewInitMetaRepo(); factoryErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize; %w", factoryErr)
	} else if thatMetaRepo, initErr := InitThatMetaRepo(); initErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize path to meta repo; %w", initErr)
	} else if runErr := initCmd.Run(thatMetaRepo); runErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize repository; %w", runErr)
	} else {
		return nil
	}
}

// Set a standardized path to a meta repo in the local test directory, adding a hook to clear it.
func InitThatMetaRepo() (string, error) {
	if _thatMetaRepo != "" {
		return "", fmt.Errorf("meta_repo_helper: meta repo path already configured: %s", _thatMetaRepo)
	} else if testDir, testDirErr := TestDir(); testDirErr != nil {
		return "", fmt.Errorf("meta_repo_helper: failed to access test directory; %w", testDirErr)
	} else {
		setThatMetaRepo(filepath.Join(testDir, "meta"))
		return _thatMetaRepo, nil
	}
}

func setThatMetaRepo(path string) {
	_thatMetaRepo = path
}

// A path to a meta repo which has been initialized earlier in this scenario.
func ThatMetaRepo() (string, error) {
	if _thatMetaRepo == "" {
		return "", fmt.Errorf("meta_repo_helper: no meta repo has been configured")
	}

	return _thatMetaRepo, nil
}
