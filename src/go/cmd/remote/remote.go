package cmdremote

import (
	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewRemoteCmd(
	group cmdshared.CommandGroup,
	parser cmdshared.CliConfigParser,
) *cobra.Command {
	//TODO KDK: Add run command that shows debug if requested or shows help if no further arguments are given
	remoteCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: group.Id,
		Long:    "Deal with repositories on remote hosts.",
		Short:   "Deal with remote repositories",
		Use:     "remote",
	}

	remoteCmd.AddCommand(NewListRemoteCmd(parser))
	remoteCmd.AddCommand(NewRegisterRemoteCmd(parser))
	return remoteCmd
}
