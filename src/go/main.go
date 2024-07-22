package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kkrull/marmot/cmd"
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
	cliFactory := cmd.NewCliFactory().WithStdIO(stdout, stderr)
	if configErr := cliFactory.ForExecutable(); configErr != nil {
		return configErr
	} else if rootCmd, buildErr := cliFactory.CommandTree(); buildErr != nil {
		return buildErr
	} else if executeErr := rootCmd.Execute(); executeErr != nil {
		return executeErr
	} else {
		return nil
	}
}
