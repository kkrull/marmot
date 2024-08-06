package cmd

import (
	cmdremote "github.com/kkrull/marmot/cmd/remote"
	cmdroot "github.com/kkrull/marmot/cmd/root"
	"github.com/spf13/cobra"
)

// Construct a CLI command to deal with remote repositories.
func NewRemoteCommand() *remoteCommand {
	return &remoteCommand{}
}

type remoteCommand struct{}

// Map to a command that runs on Cobra.
func (cliCmd *remoteCommand) ToCobraCommand() *cobra.Command {
	remoteCobraCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: repositoryGroup.id(),
		Long:    "Deal with repositories on remote hosts.",
		RunE:    runRemote,
		Short:   "Deal with remote repositories",
		Use:     "remote",
	}

	remoteCobraCmd.AddCommand(
		cmdremote.NewListCommand().ToCobraCommand(),
		cmdremote.NewRegisterCommand().ToCobraCommand(),
	)
	return remoteCobraCmd
}

func runRemote(cli *cobra.Command, args []string) error {
	if parser, newErr := cmdroot.RootCommandParser(); newErr != nil {
		return newErr
	} else if config, parseErr := parser.Parse(cli.Flags(), args); parseErr != nil {
		return parseErr
	} else if config.Debug() {
		config.PrintDebug(cli.OutOrStdout())
		return nil
	} else if len(args) == 0 {
		return cli.Help()
	} else {
		return nil
	}
}
