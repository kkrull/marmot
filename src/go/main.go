package main

import (
	"fmt"
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
		fmt.Fprintln(stderr, err.Error())
		os.Exit(1)
	}
}

func mainE() error {
	if version, readErr := core.ExecutableVersion(); readErr != nil {
		return readErr
	} else if initErr := core.SetMarmotVersion(version); initErr != nil {
		return initErr
	} else {
		cmd := cmd.NewRootCmd(version)
		return cmd.Execute()
	}
}
