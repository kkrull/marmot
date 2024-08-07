package cukesupport

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/cucumber/godog"
)

/* Hook */

// The name of the tag to identify in feature files.
const MetaRepoTag = "@MetaRepo"

// Add hook and fixtures to reset state between scenarios.  Depends on @LocalDir, so run after that.
func addMetaRepoFixtureAfterLocalDir(ctx *godog.ScenarioContext) {
	ctx.After(afterMetaRepo)
	ctx.Before(beforeMetaRepo)
}

func afterMetaRepo(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
	//Always clear state; it could have been initialized by a hook _or_ an explicit step
	forgetThatMetaRepo()
	return ctx, err
}

func beforeMetaRepo(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	if hookTag := findTag(MetaRepoTag, sc.Tags); hookTag == nil {
		return ctx, nil
	} else {
		return ctx, InitNewMetaRepo("42")
	}
}

/* State */

var _thatMetaRepo string

func forgetThatMetaRepo() {
	_thatMetaRepo = ""
}

// Set a standardized path to a meta repo in the local test directory.
func initMetaRepoPath() (string, error) {
	if _thatMetaRepo != "" {
		return "", fmt.Errorf("meta_repo_helper: meta repo path already configured: %s", _thatMetaRepo)
	} else if testDir, testDirErr := TestDir(); testDirErr != nil {
		return "", fmt.Errorf("meta_repo_helper: failed to access test directory; %w", testDirErr)
	} else {
		_thatMetaRepo = filepath.Join(testDir, "meta")
		return _thatMetaRepo, nil
	}
}

// A path to a meta repo which has been initialized earlier in this scenario.
func ThatMetaRepo() (string, error) {
	if _thatMetaRepo == "" {
		return "", fmt.Errorf("meta_repo_helper: no meta repo has been configured")
	}

	return _thatMetaRepo, nil
}

/* Initialization */

// Convenience method to initialize a new meta repo with a specified version.  Needs @LocalDir.
func InitNewMetaRepo(version string) error {
	factory := ThatCommandFactoryS(version)
	if initCmd, factoryErr := factory.NewInitMetaRepo(); factoryErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize; %w", factoryErr)
	} else if thatMetaRepo, initErr := initMetaRepoPath(); initErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize path to meta repo; %w", initErr)
	} else if runErr := initCmd.Run(thatMetaRepo); runErr != nil {
		return fmt.Errorf("meta_repo_steps: failed to initialize repository; %w", runErr)
	} else {
		return nil
	}
}
