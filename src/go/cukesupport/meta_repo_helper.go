package cukesupport

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/cucumber/godog"
)

// State to clear between scenarios
var thatMetaRepo string

// Reset the meta repo so another scenario can make its own
func clearThatMetaRepo() {
	thatMetaRepo = ""
}

// Set a standardized path to a meta repo in the local test directory, adding a hook to clear it
func InitThatMetaRepo(ctx *godog.ScenarioContext) (string, error) {
	if existing := peekThatMetaRepo(); existing != "" {
		return "", fmt.Errorf("meta_repo_helper: meta repo path already configured: %s", existing)
	} else if testDir, testDirErr := TestDir(); testDirErr != nil {
		return "", fmt.Errorf("meta_repo_helper: failed to access test directory; %w", testDirErr)
	} else {
		setThatMetaRepo(ctx, filepath.Join(testDir, "meta"))
		return peekThatMetaRepo(), nil
	}
}

func peekThatMetaRepo() string {
	return thatMetaRepo
}

func setThatMetaRepo(ctx *godog.ScenarioContext, path string) {
	thatMetaRepo = path
	ctx.After(func(ctx context.Context, _ *godog.Scenario, err error) (context.Context, error) {
		clearThatMetaRepo()
		return ctx, err
	})
}

// A path to a meta repo which has been set earlier in this scenario, or an error.
func ThatMetaRepo() (string, error) {
	if thatMetaRepo == "" {
		return "", fmt.Errorf("meta_repo_helper: no meta repo has been configured")
	}

	return thatMetaRepo, nil
}
