package cmdremote

import (
	"fmt"

	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewRemoteCmd() *cobra.Command {
	return remoteCmd
}

var remoteCmd = &cobra.Command{
	Args:    cobra.NoArgs,
	GroupID: cmdshared.RepositoryGroup.Id(),
	Long:    "Deal with repositories on remote hosts.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remote called")
	},
	Short: "Deal with remote repositories",
	Use:   "remote",
}

func init() {}
