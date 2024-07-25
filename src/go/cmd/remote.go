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

func runRemote(cobraCmd *cobra.Command, args []string) error {
	if flags, flagErr := cmdroot.RootFlagSet(); flagErr != nil {
		return flagErr
	} else if config, parseErr := flags.ParseAppConfig(cobraCmd.Flags(), args); parseErr != nil {
		return parseErr
	} else if config.Debug() {
		config.PrintDebug(cobraCmd.OutOrStdout())
		return nil
	} else if len(args) == 0 {
		return cobraCmd.Help()
	} else {
		return nil
	}
}
