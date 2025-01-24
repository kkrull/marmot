package cmdlocal

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewListLocalCmd() *cobra.Command {
	listLocalCmd := &cobra.Command{
		Args: cobra.NoArgs,
		Long: "List local repositories registered with Marmot.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("local list called")
		},
		Short: "List local repositories",
		Use:   "list",
	}

	return listLocalCmd
}
