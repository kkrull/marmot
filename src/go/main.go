package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kkrull/marmot/cmd"
)

func main() {
	var (
		stdout     io.Writer       = os.Stdout
		stderr     io.Writer       = os.Stderr
		cliFactory *cmd.CliFactory = cmd.NewCliFactory().WithStdIO(stdout, stderr)
	)

	if err := doMain(cliFactory); err != nil {
		fmt.Fprintln(stderr, err.Error())
		os.Exit(1)
	}
}

func doMain(cliFactory *cmd.CliFactory) error {
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
