package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	GroupID: metaRepoGroup,
	Long:    "Initialize a blank Meta Repo in the configured directory, if none is already present.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
	Short: "initialize a meta repo",
	Use:   "init",
}

func init() {
	rootCmd.AddCommand(initCmd)
}
