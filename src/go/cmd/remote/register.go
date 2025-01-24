package cmdremote

import (
	"io"

	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewRegisterRemoteCmd(parser cmdshared.CliConfigParser) *cobra.Command {
	registerRemoteCmd := &cobra.Command{
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Example: `marmot remote register https://github.com/drwily/skull-fortress
gh repo list --json url | jq -r '.[].url' | marmot remote register -`,
		Long: `Register Git repositories on remote hosts at each [URL].
When URL is -, stop processing arguments and read newline-delimited URLs from standard input.`,
		RunE: cmdshared.CobraCommandAdapter(
			parser.ParseC,
			makeRegisterRemotesFn(),
		),
		Short: "Register remote repositories",
		Use:   "register [flags] [URL]... [-]",
	}

	return registerRemoteCmd
}

type registerRemotesAction = func(cmdshared.CliConfig, io.Writer) error

func makeRegisterRemotesFn() registerRemotesAction {
	return func(config cmdshared.CliConfig, stdout io.Writer) error {
		if config.Debug() {
			config.PrintDebug(stdout)
			return nil
		} else if action, actErr := config.ActionFactory().NewRegisterRemoteRepositories(); actErr != nil {
			return actErr
		} else if urlsFromArgs, argErr := config.ArgsAsUrls(); argErr != nil {
			return argErr
		} else if urlsFromInput, stdInErr := config.InputLinesAsUrls(); stdInErr != nil {
			return stdInErr
		} else {
			return action.Run(append(urlsFromArgs, urlsFromInput...))
		}
	}
}
