package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// Construct a CLI command to deal with remote repositories.
func NewRemoteCommand() *remoteCommand {
	return &remoteCommand{}
}

type remoteCommand struct{}

func (cliCmd *remoteCommand) toCobraCommand() *cobra.Command {
	return &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: repositoryGroup.id(),
		Long:    "Deal with repositories on remote hosts.",
		RunE:    runRemote,
		Short:   "Deal with remote repositories",
		Use:     "remote",
	}
}

func runRemote(cobraCmd *cobra.Command, args []string) error {
	if config, parseErr := RootFlagSet().ParseAppConfig(cobraCmd.Flags(), args); parseErr != nil {
		return parseErr
	} else if config.Debug() {
		config.PrintDebug(cobraCmd.OutOrStdout())
		return nil
	} else {
		return errors.ErrUnsupported
	}
}
