package cmdlocal

import (
	"github.com/spf13/cobra"
)

func NewLocalCmd(groupId string) *cobra.Command {
	localCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: groupId,
		Long:    "Deal with repositories on the local filesystem.",
		Short:   "Deal with local repositories",
		Use:     "local",
	}

	localCmd.AddCommand(NewListLocalCmd()) // TODO KDK: Wire up and run this command
	localCmd.AddCommand(NewRegisterLocalCmd())
	return localCmd
}
