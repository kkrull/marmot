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

/* Mapping to Cobra */

// Add this command as a sub-command of the given Cobra command.
func (cliCmd *remoteCommand) AddToCobra(cobraCmd *cobra.Command) {
	cobraCmd.AddCommand(cliCmd.toCobraCommand())
}

func (cliCmd *remoteCommand) toCobraCommand() *cobra.Command {
	remoteCobraCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: repositoryGroup.id(),
		Long:    "Deal with repositories on remote hosts.",
		RunE:    runRemoteCobra,
		Short:   "Deal with remote repositories",
		Use:     "remote",
	}

	remoteCobraCmd.AddCommand(
		cmdremote.NewListCommand().ToCobraCommand(),
		cmdremote.NewRegisterCommand().ToCobraCommand(),
	)
	return remoteCobraCmd
}

func runRemoteCobra(cli *cobra.Command, args []string) error {
	if parser, newErr := cmdroot.RootConfigParser(); newErr != nil {
		return newErr
	} else if config, parseErr := parser.Parse(cli.Flags(), args); parseErr != nil {
		return parseErr
	} else if config.Debug() {
		config.PrintDebug(cli.OutOrStdout())
		return nil
	} else if len(args) == 0 {
		return cli.Help()
	} else {
		// Run the sub-command named in the arguments
		return nil
	}
}
