package cmdremote

import (
	"fmt"
	"net/url"

	cmdroot "github.com/kkrull/marmot/cmd/root"
	"github.com/spf13/cobra"
)

// Construct a CLI command to register remote repositories.
func NewRegisterCommand() *registerCommand {
	return &registerCommand{}
}

// TODO KDK: Read from stdin/pipe
type registerCommand struct{}

func (registerCommand) ToCobraCommand() *cobra.Command {
	return &cobra.Command{
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Example: `marmot remote register https://github.com/drwily/skull-fortress
gh repo list --json url | jq -r '.[].url' | marmot remote register -`,
		Long: `Register Git repositories on remote hosts at each [URL].
When URL is -, stop processing arguments and read newline-delimited URLs from standard input.`,
		RunE:  runRegister,
		Short: "Register remote repositories",
		Use:   "register [flags] [URL]... [-]",
	}
}

func anyNotUrl(args []string) error {
	for _, arg := range args {
		if urlArg, parseErr := url.Parse(arg); parseErr != nil {
			return fmt.Errorf("URL expected: %s; %w", arg, parseErr)
		} else if !urlArg.IsAbs() {
			return fmt.Errorf("absolute URL expected: %s", arg)
		}
	}

	return nil
}

func runRegister(cobraCmd *cobra.Command, args []string) error {
	if parser, newErr := cmdroot.RootCommandParser(); newErr != nil {
		return newErr
	} else if config, parseErr := parser.ParseR(cobraCmd.Flags(), args, cobraCmd.InOrStdin()); parseErr != nil {
		return parseErr
	} else if config.Debug() {
		config.PrintDebug(cobraCmd.OutOrStdout())
		return nil
	} else if argErr := anyNotUrl(config.Args()); argErr != nil {
		return argErr
	} else if stdInErr := anyNotUrl(config.InputLines()); stdInErr != nil {
	} else {
		return runRegisterAppCmd(config)
	}
}

func runRegisterAppCmd(config cmdroot.AppConfig) error {
	if appCmd, appErr := config.AppFactory().RegisterRemoteRepositoriesCommand(); appErr != nil {
		return appErr
	} else if urlsFromArgs, argErr := config.ArgsAsUrls(); argErr != nil {
		return argErr
	} else if urlsFromInput, stdInErr := config.InputLinesAsUrls(); stdInErr != nil {
		return stdInErr
	} else if runErr := appCmd.Run(append(urlsFromArgs, urlsFromInput...)); runErr != nil {
		return runErr
	} else {
		return nil
	}
}
