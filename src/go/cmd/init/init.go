package cmdinit

import (
	"fmt"
	"io"

	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewInitCmd(
	group cmdshared.CommandGroup,
	parser cmdshared.CliConfigParser,
) *cobra.Command {
	initCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: group.Id,
		Long:    "Initialize a new Meta Repo, if none is already present.",
		RunE: parseThen(
			func(cmd *cobra.Command, args []string) (cmdshared.CliConfig, error) {
				return parser.Parse(cmd.Flags(), args)
			},
			runInitAction(),
		),
		Short: "Initialize a meta repo",
		Use:   "init",
	}

	return initCmd
}

func runInitAction() useFn {
	return func(config cmdshared.CliConfig, stdout io.Writer) error {
		if config.Debug() {
			config.PrintDebug(stdout)
			return nil
		} else if action, actErr := config.ActionFactory().NewInitMetaRepo(); actErr != nil {
			return actErr
		} else if runErr := action.Run(config.MetaRepoPath()); runErr != nil {
			return runErr
		} else {
			fmt.Fprintf(stdout, "Initialized meta repo at %s\n", config.MetaRepoPath())
			return nil
		}
	}
}

/* General purpose */

type cobraRunEFn = func(cli *cobra.Command, args []string) error

func parseThen[TConfig any](
	parseConfig func(*cobra.Command, []string) (TConfig, error),
	useConfig func(TConfig, io.Writer) error,
) cobraRunEFn {
	return func(cli *cobra.Command, args []string) error {
		if config, parseErr := parseConfig(cli, args); parseErr != nil {
			return parseErr
		} else if useErr := useConfig(config, cli.OutOrStdout()); useErr != nil {
			return useErr
		} else {
			return nil
		}
	}
}

type useFn = func(cmdshared.CliConfig, io.Writer) error
