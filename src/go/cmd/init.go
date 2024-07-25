package cmd

import (
	"fmt"

	cmdroot "github.com/kkrull/marmot/cmd/root"
	"github.com/spf13/cobra"
)

// Construct a CLI command to initialize a meta repo.
func NewInitCommand() *initCommand {
	return &initCommand{}
}

type initCommand struct{}

func (cliCmd *initCommand) ToCobraCommand() *cobra.Command {
	return &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: metaRepoGroup.id(),
		Long:    "Initialize a new Meta Repo, if none is already present.",
		RunE:    runInit,
		Short:   "Initialize a meta repo",
		Use:     "init",
	}
}

func runInit(cobraCmd *cobra.Command, args []string) error {
	if flags, flagErr := cmdroot.RootFlagSet(); flagErr != nil {
		return flagErr
	} else if config, parseErr := flags.ParseAppConfig(cobraCmd.Flags(), args); parseErr != nil {
		return parseErr
	} else if config.Debug() {
		config.PrintDebug(cobraCmd.OutOrStdout())
		return nil
	} else {
		return runInitAppCmd(cobraCmd, config)
	}
}

func runInitAppCmd(cobraCmd *cobra.Command, config cmdroot.AppConfig) error {
	if initAppCmd, initErr := config.AppFactory().InitCommand(); initErr != nil {
		return initErr
	} else if runErr := initAppCmd.Run(config.MetaRepoPath()); runErr != nil {
		return runErr
	} else {
		fmt.Fprintf(cobraCmd.OutOrStdout(), "Initialized meta repo at %s\n", config.MetaRepoPath())
		return nil
	}
}
