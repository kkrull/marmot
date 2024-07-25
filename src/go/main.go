package main

import (
	"io"
	"os"

	"github.com/kkrull/marmot/cmd"
)

var (
	stdout io.Writer = os.Stdout
	stderr io.Writer = os.Stderr
)

func main() {
	if err := mainE(); err != nil {
		os.Exit(1)
	}
}

func mainE() error {
	if cliFactory, cliErr := newCliFactory(); cliErr != nil {
		return cliErr
	} else {
		rootCmd := cliFactory.ToRootCobraCommand()
		return rootCmd.Execute()
	}
}

func newCliFactory() (*cmd.CliFactory, error) {
	return cmd.
		NewCliFactory().
		WithStdIO(stdout, stderr).
		ForExecutable()
}
