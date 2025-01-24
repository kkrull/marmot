package cmdremote

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewListRemoteCmd() *cobra.Command {
	listRemoteCmd := &cobra.Command{
		Args: cobra.NoArgs,
		Long: "List remote repositories registered with Marmot.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("remote list called")
		},
		Short: "List remote repositories",
		Use:   "list",
	}

	return listRemoteCmd
}
