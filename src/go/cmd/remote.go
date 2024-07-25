package cmd

import (
	cmdcore "github.com/kkrull/marmot/cmd/core"
	cmdremote "github.com/kkrull/marmot/cmd/remote"
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

	remoteCobraCmd.AddCommand(cmdremote.NewListCommand().ToCobraCommand())
	return remoteCobraCmd
}

func runRemote(cobraCmd *cobra.Command, args []string) error {
	if config, parseErr := cmdcore.RootFlagSet().ParseAppConfig(cobraCmd.Flags(), args); parseErr != nil {
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
