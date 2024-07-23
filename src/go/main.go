package main

import (
	"fmt"
	"io"
	"os"

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
	if appFactory, appErr := mainfactory.DefaultAppFactory(); appErr != nil {
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

func newCliFactory(appFactory *mainfactory.AppFactory) (*mainfactory.CliFactory, error) {
	return mainfactory.
		NewCliFactory(appFactory).
		WithStdIO(stdout, stderr).
		ForExecutable()
}
