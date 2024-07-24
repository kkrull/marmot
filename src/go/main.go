package main

import (
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
		os.Exit(1)
	}
}

func doMain() error {
	if cliFactory, cliErr := newCliFactory(); cliErr != nil {
		return cliErr
	} else if rootCmd, buildErr := cliFactory.CommandTree(); buildErr != nil {
		return buildErr
	} else if executeErr := rootCmd.Execute(); executeErr != nil {
		return executeErr
	} else {
		return nil
	}
}

func newCliFactory() (*mainfactory.CliFactory, error) {
	return mainfactory.
		NewCliFactory().
		WithStdIO(stdout, stderr).
		ForExecutable()
}
