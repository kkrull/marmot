package cmd

import (
	"fmt"
	"io"

	cmdroot "github.com/kkrull/marmot/cmd/root"
	"github.com/spf13/cobra"
)

// Construct a CLI command to initialize a meta repo.
func NewInitCommand() *initCommand {
	return &initCommand{}
}

type initCommand struct{}

// Map to a command that runs on Cobra.
func (cliCmd *initCommand) ToCobraCommand() *cobra.Command {
	return &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: metaRepoGroup.id(),
		Long:    "Initialize a new Meta Repo, if none is already present.",
		RunE:    runInitCobra,
		Short:   "Initialize a meta repo",
		Use:     "init",
	}
}

func runInitCobra(cli *cobra.Command, args []string) error {
	if parser, newErr := cmdroot.RootConfigParser(); newErr != nil {
		return newErr
	} else if config, parseErr := parser.Parse(cli.Flags(), args); parseErr != nil {
		return parseErr
	} else if config.Debug() {
		config.PrintDebug(cli.OutOrStdout())
		return nil
	} else {
		return runInit(config, cli.OutOrStdout())
	}
}

func runInit(config cmdroot.CliConfig, stdout io.Writer) error {
	if appCmd, initErr := config.CommandFactory().NewInitMetaRepo(); initErr != nil {
		return initErr
	} else if runErr := appCmd.Run(config.MetaRepoPath()); runErr != nil {
		return runErr
	} else {
		fmt.Fprintf(stdout, "Initialized meta repo at %s\n", config.MetaRepoPath())
		return nil
	}
}
