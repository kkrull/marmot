package cmdlocal

import (
	"io"

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
		RunE: cmdshared.CobraCommandAdapter(
			parser.ParseC,
			makeRegisterLocalsFn(),
		),
		Short: "Register local repositories",
		Use:   "register [flags] [URL]... [-]",
	}

	return registerLocalCmd
}

type registerLocalsAction = func(cmdshared.CliConfig, io.Writer) error

func makeRegisterLocalsFn() registerLocalsAction {
	return func(config cmdshared.CliConfig, stdout io.Writer) error {
		if config.Debug() {
			config.PrintDebug(stdout)
			return nil
		} else if action, actErr := config.ActionFactory().NewRegisterLocalRepositories(); actErr != nil {
			return actErr
		} else {
			argsThenInputLines := config.ArgsTrimmed()
			argsThenInputLines = append(argsThenInputLines, config.InputLinesTrimmed()...)
			return action.Run(argsThenInputLines)
		}
	}
}
