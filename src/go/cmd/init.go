package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Construct a CLI command to initialize a meta repo.
func NewInitCommand() *initCommand {
	return &initCommand{}
}

type initCommand struct{}

func (cliCmd *initCommand) RegisterWithCobra(parentCmd *cobra.Command) {
	cobraCmd := cliCmd.toCobraCommand()
	AddMetaRepoCommand(parentCmd, *cobraCmd)
}

func (cliCmd *initCommand) toCobraCommand() *cobra.Command {
	return &cobra.Command{
		Args:  cobra.NoArgs,
		Long:  "Initialize a new Meta Repo, if none is already present.",
		RunE:  runInit,
		Short: "Initialize a meta repo",
		Use:   "init",
	}
}

func runInit(cobraCmd *cobra.Command, _args []string) error {
	if config, parseErr := ParseFlags(cobraCmd); parseErr != nil {
		return parseErr
	} else if config.Debug() {
		config.PrintDebug(cobraCmd.OutOrStdout())
		return nil
	} else {
		return runInitAppCmd(cobraCmd, config)
	}
}

func runInitAppCmd(cobraCmd *cobra.Command, config AppConfig) error {
	initAppCmd := config.AppFactory().InitCommand()
	if runErr := initAppCmd.Run(config.MetaRepoPath()); runErr != nil {
		return runErr
	} else {
		fmt.Fprintf(cobraCmd.OutOrStdout(), "Initialized meta repo at %s\n", config.MetaRepoPath())
		return nil
	}
}
