package cmdlocal

import (
	"fmt"

	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewListLocalCmd(parser cmdshared.CliConfigParser) *cobra.Command {
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
