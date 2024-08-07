package cukesupport

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/cucumber/godog"
)

// State to clear between scenarios
var _thatMetaRepo string

// Add fixtures to reset state between scenarios.
func addMetaRepoFixture(ctx *godog.ScenarioContext) {
	ctx.After(func(ctx context.Context, _ *godog.Scenario, err error) (context.Context, error) {
		clearThatMetaRepo()
		return ctx, err
	})
}

// Reset the meta repo so another scenario can make its own.
func clearThatMetaRepo() {
	_thatMetaRepo = ""
}

// Set a standardized path to a meta repo in the local test directory, adding a hook to clear it.
func InitThatMetaRepo() (string, error) {
	if existing := peekThatMetaRepo(); existing != "" {
		return "", fmt.Errorf("meta_repo_helper: meta repo path already configured: %s", existing)
	} else if testDir, testDirErr := TestDir(); testDirErr != nil {
		return "", fmt.Errorf("meta_repo_helper: failed to access test directory; %w", testDirErr)
	} else {
		setThatMetaRepo(filepath.Join(testDir, "meta"))
		return peekThatMetaRepo(), nil
	}
}

func peekThatMetaRepo() string {
	return _thatMetaRepo
}

func setThatMetaRepo(path string) {
	_thatMetaRepo = path
}

// A path to a meta repo which has been set earlier in this scenario.
func ThatMetaRepo() (string, error) {
	if _thatMetaRepo == "" {
		return "", fmt.Errorf("meta_repo_helper: no meta repo has been configured")
	}

	return _thatMetaRepo, nil
}
