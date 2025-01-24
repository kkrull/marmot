package cmd

import (
	"fmt"

	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Args:    cobra.NoArgs,
	GroupID: cmdshared.MetaRepoGroup.Id(),
	Long:    "Initialize a new Meta Repo, if none is already present.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
	Short: "Initialize a meta repo",
	Use:   "init",
}

func init() {}
