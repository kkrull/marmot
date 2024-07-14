package cukesupport

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
)

// State to clear between scenarios
var thatMetaRepo string

// Reset the meta repo so another scenario can make its own
func clearThatMetaRepo() {
	thatMetaRepo = ""
}

// A path to a meta repo which has been set earlier in this scenario, or an empty string.
func PeekThatMetaRepo() string {
	return thatMetaRepo
}

// Save the path to the meta repo for later use, adding an after hook that clears state afterwards
func SetThatMetaRepo(ctx *godog.ScenarioContext, path string) string {
	thatMetaRepo = path
	ctx.After(func(ctx context.Context, _ *godog.Scenario, err error) (context.Context, error) {
		clearThatMetaRepo()
		return ctx, err
	})

	return thatMetaRepo
}

// A path to a meta repo which has been set earlier in this scenario, or an error.
func ThatMetaRepo() (string, error) {
	if thatMetaRepo == "" {
		return "", fmt.Errorf("meta_repo_steps: no meta repo has been configured")
	}

	return thatMetaRepo, nil
}
