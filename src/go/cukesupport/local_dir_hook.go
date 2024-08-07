package cukesupport

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/cucumber/godog"
)

// The name of the tag to identify in feature files
const LocalDirTag = "@LocalDir"

/* State */

var _testDir string

// A temporary directory for scenarios tagged with `TagName`, to automatically delete afterwards.
func TestDir() (string, error) {
	if _testDir == "" {
		return "", fmt.Errorf("[%s] not initialized", LocalDirTag)
	}

	return _testDir, nil
}

// The named subdirectory inside this scenario's temporary directory (which is deleted afterwards).
func TestSubDir(subdir string) (string, error) {
	if _testDir == "" {
		return "", fmt.Errorf("[%s] not initialized", LocalDirTag)
	}

	return filepath.Join(_testDir, subdir), nil
}

func setTestDir(path string) {
	_testDir = path
}

/* Hooks */

// Make this hook available to a scenario, if it is tagged with its name.
func addLocalDirHook(ctx *godog.ScenarioContext) {
	ctx.After(afterLocalDir)
	ctx.Before(beforeLocalDir)
}

func afterLocalDir(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
	if _testDir == "" || findTag(LocalDirTag, sc.Tags) == nil {
		return ctx, err
	} else if rmErr := os.RemoveAll(_testDir); rmErr != nil {
		hookErr := fmt.Errorf("%s: failed to remove %s; %w", LocalDirTag, _testDir, rmErr)
		return ctx, errors.Join(err, hookErr)
	} else {
		log.Printf("[%s] Removed test directory: %s\n", LocalDirTag, _testDir)
		setTestDir("")
		return ctx, err
	}
}

func beforeLocalDir(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	if hookTag := findTag(LocalDirTag, sc.Tags); hookTag == nil {
		return ctx, nil
	} else if localDir, pathErr := filepath.Abs(filepath.Join(".", "localDir")); pathErr != nil {
		return ctx, fmt.Errorf("%s: failed to determine path; %w", LocalDirTag, pathErr)
	} else if mkErr := os.MkdirAll(localDir, 0o777); mkErr != nil {
		return ctx, fmt.Errorf("%s: failed to create %s; %w", LocalDirTag, localDir, mkErr)
	} else {
		log.Printf("[%s] Created test directory: %s\n", LocalDirTag, localDir)
		setTestDir(localDir)
		return ctx, nil
	}
}
