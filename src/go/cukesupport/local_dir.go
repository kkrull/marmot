package cukesupport

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/cucumber/godog"
	messages "github.com/cucumber/messages/go/v21"
)

// The name of the tag to identify in feature files
const TagName = "@LocalDir"

/* State */

var testDir string

// A temporary directory created for the current scenario, which will be deleted later.  Tag
// scenarios or features with the tag in `TagName` to get started.
func TestDir() (string, error) {
	if testDir == "" {
		return "", fmt.Errorf("[%s] not initialized", TagName)
	}

	return testDir, nil
}

func setTestDir(path string) {
	testDir = path
}

/* Hooks */

// Add hooks for this tag so that it runs on matching scenarios
func AddTo(ctx *godog.ScenarioContext) {
	ctx.After(afterHook)
	ctx.Before(beforeHook)
}

func afterHook(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
	matchingTag := findTag(TagName, sc.Tags)
	if nothingToDo := testDir == ""; nothingToDo {
		return ctx, err
	} else if notApplicable := matchingTag == nil; notApplicable {
		return ctx, err
	} else if rmErr := os.RemoveAll(testDir); rmErr != nil {
		return ctx, errors.Join(err, fmt.Errorf("%s: failed to remove %s; %w", TagName, testDir, rmErr))
	} else {
		log.Printf("[%s] Removed test directory: %s\n", TagName, testDir)
		setTestDir("")
		return ctx, err
	}
}

func beforeHook(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	if hookTag := findTag(TagName, sc.Tags); hookTag == nil {
		return ctx, nil
	} else if localDir, pathErr := filepath.Abs(filepath.Join(".", "localDir")); pathErr != nil {
		return ctx, fmt.Errorf("%s: failed to determine path; %w", TagName, pathErr)
	} else if mkErr := os.MkdirAll(localDir, 0o777); mkErr != nil {
		return ctx, fmt.Errorf("%s: failed to create %s; %w", TagName, localDir, mkErr)
	} else {
		log.Printf("[%s] Created test directory: %s\n", TagName, localDir)
		setTestDir(localDir)
		return ctx, nil
	}
}

func findTag(name string, tags []*messages.PickleTag) *messages.PickleTag {
	for _, tag := range tags {
		if tag.Name == name {
			return tag
		}
	}

	return nil
}
