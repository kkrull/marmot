package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kkrull/marmot/mainfactory"
)

var (
	stdout io.Writer = os.Stdout
	stderr io.Writer = os.Stderr
)

func main() {
	if err := doMain(); err != nil {
		fmt.Fprintln(stderr, err.Error())
		os.Exit(1)
	}
}

func doMain() error {
	if appFactory, appErr := defaultAppFactory(); appErr != nil {
		return appErr
	} else if cliFactory, cliErr := newCliFactory(appFactory); cliErr != nil {
		return cliErr
	} else if rootCmd, buildErr := cliFactory.CommandTree(); buildErr != nil {
		return buildErr
	} else if executeErr := rootCmd.Execute(); executeErr != nil {
		return executeErr
	} else {
		return nil
	}
}

/* App factory */

func defaultAppFactory() (*mainfactory.AppFactory, error) {
	if metaRepoPath, pathErr := defaultMetaRepoPath(); pathErr != nil {
		return nil, pathErr
	} else {
		return newAppFactory().ForLocalMetaRepo(metaRepoPath), nil
	}
}

func defaultMetaRepoPath() (string, error) {
	if homeDir, homeErr := os.UserHomeDir(); homeErr != nil {
		return "", fmt.Errorf("failed to locate home directory; %w", homeErr)
	} else {
		return filepath.Join(homeDir, "meta"), nil
	}
}

func newAppFactory() *mainfactory.AppFactory {
	return &mainfactory.AppFactory{}
}

/* CLI factory */

func newCliFactory(appFactory *mainfactory.AppFactory) (*mainfactory.CliFactory, error) {
	return mainfactory.
		NewCliFactory(appFactory).
		WithStdIO(stdout, stderr).
		ForExecutable()
}
