package cmdlocal

import (
	"fmt"

	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewRegisterLocalCmd(parser cmdshared.CliConfigParser) *cobra.Command {
	registerLocalCmd := &cobra.Command{
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Example: `marmot local register ~/git/drwily.github.com/skull-fortress
find ~/git -type d -name '.git' -exec marmot local register - {} +`,
		Long: `Register Git repositories on the local filesystem at each [PATH].
When PATH is -, stop processing arguments and read newline-delimited paths from standard input.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("local register called")
		},
		Short: "Register local repositories",
		Use:   "register [flags] [URL]... [-]",
	}

	return registerLocalCmd
}
