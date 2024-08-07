package cmdremote

import (
	"fmt"
	"io"

	cmdroot "github.com/kkrull/marmot/cmd/root"
	"github.com/spf13/cobra"
)

// Construct a CLI command to list remote repositories.
func NewListCommand() *listRemoteCommand {
	return &listRemoteCommand{}
}

type listRemoteCommand struct{}

func runList(config cmdroot.CliConfig, stdout io.Writer) error {
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

/* Mapping to Cobra */

// Add this command as a sub-command of the given Cobra command.
func (cliCmd *listRemoteCommand) AddToCobra(cobraCmd *cobra.Command) {
	cobraCmd.AddCommand(cliCmd.toCobraCommand())
}

func (listRemoteCommand) toCobraCommand() *cobra.Command {
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
