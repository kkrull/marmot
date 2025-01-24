package cmdlocal

import (
	cmdroot "github.com/kkrull/marmot/cmdv1/root"
	"github.com/spf13/cobra"
)

// Construct a CLI command to register local repositories.
func NewRegisterCommand() *registerLocalCommand {
	return &registerLocalCommand{}
}

type registerLocalCommand struct{}

func runRegister(config cmdroot.CliConfig) error {
	if appCmd, appErr := config.CommandFactory().NewRegisterLocalRepositories(); appErr != nil {
		return appErr
	} else {
		argsThenInputLines := config.ArgsTrimmed()
		argsThenInputLines = append(argsThenInputLines, config.InputLinesTrimmed()...)
		return appCmd.Run(argsThenInputLines)
	}
}

/* Mapping to Cobra */

// Add this command as a sub-command of the given Cobra command.
func (cliCmd *registerLocalCommand) AddToCobra(cobraCmd *cobra.Command) {
	cobraCmd.AddCommand(cliCmd.toCobraCommand())
}

func (registerLocalCommand) toCobraCommand() *cobra.Command {
	return &cobra.Command{
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Example: `marmot local register ~/git/drwily.github.com/skull-fortress
find ~/git -type d -name '.git' -exec marmot local register - {} +`,
		Long: `Register Git repositories on the local filesystem at each [PATH].
When PATH is -, stop processing arguments and read newline-delimited paths from standard input.`,
		RunE:  runRegisterCobra,
		Short: "Register local repositories",
		Use:   "register [flags] [URL]... [-]",
	}
}

func runRegisterCobra(cli *cobra.Command, args []string) error {
	if parser, newErr := cmdroot.RootConfigParser(); newErr != nil {
		return newErr
	} else if config, parseErr := parser.ParseR(cli.Flags(), args, cli.InOrStdin()); parseErr != nil {
		return parseErr
	} else if config.Debug() {
		config.PrintDebug(cli.OutOrStdout())
		return nil
	} else {
		return runRegister(config)
	}
}
