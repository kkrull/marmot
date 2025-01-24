package cmdinit

import (
	"fmt"

	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewInitCmd(group cmdshared.CommandGroup, parser cmdshared.CliConfigParser) *cobra.Command {
	initCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: group.Id,
		Long:    "Initialize a new Meta Repo, if none is already present.",
		RunE:    newCobraCommandRunE(parser),
		Short:   "Initialize a meta repo",
		Use:     "init",
	}

	return initCmd
}

type cobraRunner = func(cmd *cobra.Command, args []string) error

func newCobraCommandRunE(parser cmdshared.CliConfigParser) cobraRunner {
	return func(cli *cobra.Command, args []string) error {
		if config, parseErr := parser.Parse(cli.Flags(), args); parseErr != nil {
			return parseErr
		} else if config.Debug() {
			config.PrintDebug(cli.OutOrStdout())
			return nil
		} else if action, actErr := config.ActionFactory().NewInitMetaRepo(); actErr != nil {
			return actErr
		} else if runErr := action.Run(config.MetaRepoPath()); runErr != nil {
			return runErr
		} else {
			fmt.Fprintf(cli.OutOrStdout(), "Initialized meta repo at %s\n", config.MetaRepoPath())
			return nil
		}
	}
}
