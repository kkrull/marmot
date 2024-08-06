package cmdremote

import (
	"fmt"
	"io"

	cmdroot "github.com/kkrull/marmot/cmd/root"
	"github.com/spf13/cobra"
)

// Construct a CLI command to list remote repositories.
func NewListCommand() *listCommand {
	return &listCommand{}
}

type listCommand struct{}

// Map to a command that runs on Cobra.
func (listCommand) ToCobraCommand() *cobra.Command {
	return &cobra.Command{
		Args:  cobra.NoArgs,
		Long:  "List remote repositories registered with Marmot.",
		RunE:  runListCobra,
		Short: "List remote repositories",
		Use:   "list",
	}
}

func runListCobra(cli *cobra.Command, args []string) error {
	if parser, newErr := cmdroot.RootConfigParser(); newErr != nil {
		return newErr
	} else if config, parseErr := parser.Parse(cli.Flags(), args); parseErr != nil {
		return parseErr
	} else if config.Debug() {
		config.PrintDebug(cli.OutOrStdout())
		return nil
	} else {
		return runList(config, cli.OutOrStdout())
	}
}

func runList(config cmdroot.CliConfig, stdout io.Writer) error {
	queryFactory := config.QueryFactory()
	if listRemoteRepositories, appErr := queryFactory.NewListRemoteRepositories(); appErr != nil {
		return appErr
	} else if repositories, runErr := listRemoteRepositories(); runErr != nil {
		return runErr
	} else {
		for _, repository := range repositories.RemoteHrefs() {
			fmt.Fprintf(stdout, "%s\n", repository)
		}
		return nil
	}
}
