package cmdinit

import (
	"fmt"
	"io"

	cmdshared "github.com/kkrull/marmot/cmd/shared"
	cmdroot "github.com/kkrull/marmot/cmdv1/root"
	"github.com/spf13/cobra"
)

func NewInitCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: cmdshared.MetaRepoGroup.Id(),
		Long:    "Initialize a new Meta Repo, if none is already present.",
		RunE:    newCobraCommandRunE(),
		Short:   "Initialize a meta repo",
		Use:     "init",
	}

	return initCmd
}

type cobraRunner = func(cmd *cobra.Command, args []string) error

func newCobraCommandRunE() cobraRunner {
	return func(cli *cobra.Command, args []string) error {
		if parser, newErr := cmdroot.RootConfigParser(); newErr != nil {
			return newErr
		} else if config, parseErr := parser.Parse(cli.Flags(), args); parseErr != nil {
			return parseErr
		} else if config.Debug() {
			config.PrintDebug(cli.OutOrStdout())
			return nil
		} else {
			return runInit(config, cli.OutOrStdout())
		}
	}
}

func runInit(config cmdroot.CliConfig, stdout io.Writer) error {
	if action, actErr := config.ActionFactory().NewInitMetaRepo(); actErr != nil {
		return actErr
	} else if runErr := action.Run(config.MetaRepoPath()); runErr != nil {
		return runErr
	} else {
		fmt.Fprintf(stdout, "Initialized meta repo at %s\n", config.MetaRepoPath())
		return nil
	}
}
