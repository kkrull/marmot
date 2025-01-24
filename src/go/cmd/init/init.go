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
		RunE: cmdshared.CobraCommandAdapter(
			parser.ParseC,
			makeInitFn(),
		),
		Short: "Initialize a meta repo",
		Use:   "init",
	}

	return initCmd
}

type initAction = func(cmdshared.CliConfig, io.Writer) error

func makeInitFn() initAction {
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
