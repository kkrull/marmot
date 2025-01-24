package cmdlocal

import (
	"fmt"

	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewLocalCmd() *cobra.Command {
	localCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: cmdshared.RepositoryGroup.Id(),
		Long:    "Deal with repositories on the local filesystem.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("local called")
		},
		Short: "Deal with local repositories",
		Use:   "local",
	}

	return localCmd
}
