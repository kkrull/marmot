package cmdremote

import (
	"github.com/spf13/cobra"
)

func NewRemoteCmd(groupId string) *cobra.Command {
	remoteCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: groupId,
		Long:    "Deal with repositories on remote hosts.",
		Short:   "Deal with remote repositories",
		Use:     "remote",
	}

	remoteCmd.AddCommand(NewListRemoteCmd())
	remoteCmd.AddCommand(NewRegisterRemoteCmd())
	return remoteCmd
}
