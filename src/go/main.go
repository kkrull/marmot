package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kkrull/marmot/cmd"
	"github.com/kkrull/marmot/core"
)

var (
	stdout io.Writer = os.Stdout
	stderr io.Writer = os.Stderr
)

func main() {
	if err := mainE(); err != nil {
		fmt.Fprintln(stderr, err.Error())
		os.Exit(1)
	}
}

func mainE() error {
	if version, readErr := core.ExecutableVersion(); readErr != nil {
		return readErr
	} else if initErr := core.SetMarmotVersion(version); initErr != nil {
		return initErr
	} else if homeDir, homeErr := os.UserHomeDir(); homeErr != nil {
		return homeErr
	} else {
		metaRepoDefault := filepath.Join(homeDir, "meta")

		cmd := cmd.NewRootCmd(metaRepoDefault, version)
		cmd.SetOut(stdout)
		cmd.SetErr(stderr)

		return cmd.Execute()
	}
}
