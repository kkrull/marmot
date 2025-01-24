package cmdlocal

import (
	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewLocalCmd(group cmdshared.CommandGroup) *cobra.Command {
	localCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: group.Id,
		Long:    "Deal with repositories on the local filesystem.",
		Short:   "Deal with local repositories",
		Use:     "local",
	}

	localCmd.AddCommand(NewListLocalCmd()) // TODO KDK: Wire up and run this command
	localCmd.AddCommand(NewRegisterLocalCmd())
	return localCmd
}
