package cmdlocal

import (
	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewLocalCmd() *cobra.Command {
	localCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: cmdshared.RepositoryGroup.Id(),
		Long:    "Deal with repositories on the local filesystem.",
		Short: "Deal with local repositories",
		Use:   "local",
	}

	localCmd.AddCommand(NewListLocalCmd())
	localCmd.AddCommand(NewRegisterLocalCmd())
	return localCmd
}
