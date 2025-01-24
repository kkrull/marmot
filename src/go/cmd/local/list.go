package cmdlocal

import (
	"fmt"
	"io"

	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewListLocalCmd(parser cmdshared.CliConfigParser) *cobra.Command {
	listLocalCmd := &cobra.Command{
		Args: cobra.NoArgs,
		Long: "List local repositories registered with Marmot.",
		RunE: cmdshared.CobraCommandAdapter(
			parser.ParseC,
			makeListLocalsFn(),
		),
		Short: "List local repositories",
		Use:   "list",
	}

	return listLocalCmd
}

type listLocalsAction = func(cmdshared.CliConfig, io.Writer) error

func makeListLocalsFn() listLocalsAction {
	return func(config cmdshared.CliConfig, stdout io.Writer) error {
		if config.Debug() {
			config.PrintDebug(stdout)
			return nil
		}

		queryFactory := config.QueryFactory()
		if listRepositories, appErr := queryFactory.NewListLocalRepositories(); appErr != nil {
			return appErr
		} else if repositories, runErr := listRepositories(); runErr != nil {
			return runErr
		} else {
			for _, localPath := range repositories.LocalPaths() {
				fmt.Fprintf(stdout, "%s\n", localPath)
			}

			return nil
		}
	}
}
