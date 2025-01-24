package cmd

import (
	"fmt"

	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

var localCmd = &cobra.Command{
	Args:    cobra.NoArgs,
	GroupID: cmdshared.RepositoryGroup.Id(),
	Long:    "Deal with repositories on the local filesystem.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("local called")
	},
	Short: "Deal with local repositories",
	Use:   "local",
}

func init() {}
