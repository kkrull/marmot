package cmdremote

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRegisterRemoteCmd() *cobra.Command {
	registerRemoteCmd := &cobra.Command{
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Example: `marmot remote register https://github.com/drwily/skull-fortress
gh repo list --json url | jq -r '.[].url' | marmot remote register -`,
		Long: `Register Git repositories on remote hosts at each [URL].
When URL is -, stop processing arguments and read newline-delimited URLs from standard input.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("remote register called")
		},
		Short: "Register remote repositories",
		Use:   "register [flags] [URL]... [-]",
	}

	return registerRemoteCmd
}
