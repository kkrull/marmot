package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var remoteCmd = &cobra.Command{
	Args:    cobra.NoArgs,
	GroupID: repositoryGroup.id(),
	Long:    "Deal with repositories on remote hosts.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remote called")
	},
	Short: "Deal with remote repositories",
	Use:   "remote",
}

func init() {}
