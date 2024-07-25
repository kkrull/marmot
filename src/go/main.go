package main

import (
	"io"
	"os"

	"github.com/kkrull/marmot/cmd"
	"github.com/kkrull/marmot/core"
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
	if version, versionErr := core.ExecutableVersion(); versionErr != nil {
		return versionErr
	} else {
		cliFactory := cmd.
			NewCliFactory(version).
			WithStdIO(stdout, stderr)
		rootCmd := cliFactory.ToRootCobraCommand()
		return rootCmd.Execute()
	}
}
