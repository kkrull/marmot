package cmdlocal

import (
	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewLocalCmd(
	group cmdshared.CommandGroup,
	parser cmdshared.CliConfigParser,
) *cobra.Command {
	//TODO KDK: Add run command that shows debug if requested or shows help if no further arguments are given
	localCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: group.Id,
		Long:    "Deal with repositories on the local filesystem.",
		Short:   "Deal with local repositories",
		Use:     "local",
	}

	localCmd.AddCommand(NewListLocalCmd(parser))
	localCmd.AddCommand(NewRegisterLocalCmd(parser))
	return localCmd
}
