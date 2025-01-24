package cmdremote

import (
	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewRemoteCmd() *cobra.Command {
	remoteCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: cmdshared.RepositoryGroup.Id(),
		Long:    "Deal with repositories on remote hosts.",
		Short:   "Deal with remote repositories",
		Use:     "remote",
	}

	remoteCmd.AddCommand(NewListRemoteCmd())
	remoteCmd.AddCommand(NewRegisterRemoteCmd())
	return remoteCmd
}
