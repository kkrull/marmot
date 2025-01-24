package cmdremote

import (
	"fmt"
	"io"

	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewListRemoteCmd(parser cmdshared.CliConfigParser) *cobra.Command {
	listRemoteCmd := &cobra.Command{
		Args: cobra.NoArgs,
		Long: "List remote repositories registered with Marmot.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("remote list called")
		},
		RunE: cmdshared.CobraCommandAdapter(
			parser.ParseC,
			makeListRemotesFn(),
		),
		Short: "List remote repositories",
		Use:   "list",
	}

	return listRemoteCmd
}

type listRemotesAction = func(cmdshared.CliConfig, io.Writer) error

func makeListRemotesFn() listRemotesAction {
	return func(config cmdshared.CliConfig, stdout io.Writer) error {
		if config.Debug() {
			config.PrintDebug(stdout)
			return nil
		}

		queryFactory := config.QueryFactory()
		if listRepositories, appErr := queryFactory.NewListRemoteRepositories(); appErr != nil {
			return appErr
		} else if repositories, runErr := listRepositories(); runErr != nil {
			return runErr
		} else {
			for _, remoteHref := range repositories.RemoteHrefs() {
				fmt.Fprintf(stdout, "%s\n", remoteHref)
			}

			return nil
		}
	}
}
