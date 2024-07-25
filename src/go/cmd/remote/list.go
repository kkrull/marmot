package cmdremote

import (
	"fmt"

	"github.com/kkrull/marmot/cmd"
	"github.com/spf13/cobra"
)

func NewListCommand() *listCommand {
	return &listCommand{}
}

type listCommand struct{}

func (listCommand) ToCobraCommand() *cobra.Command {
	return &cobra.Command{
		Args:  cobra.NoArgs,
		Long:  "List remote repositories registered with Marmot.",
		RunE:  runList,
		Short: "List remote repositories",
		Use:   "list",
	}
}

func runList(cobraCmd *cobra.Command, args []string) error {
	if config, parseErr := cmd.RootFlagSet().ParseAppConfig(cobraCmd.Flags(), args); parseErr != nil {
		return parseErr
	} else if config.Debug() {
		config.PrintDebug(cobraCmd.OutOrStdout())
		return nil
	} else {
		return runListAppCmd(cobraCmd, config)
	}
}

func runListAppCmd(cobraCmd *cobra.Command, config cmd.AppConfig) error {
	if listAppCmd, factoryErr := config.AppFactory().ListRemoteRepositoriesQuery(); factoryErr != nil {
		return factoryErr
	} else if repositories, runErr := listAppCmd(); runErr != nil {
		return runErr
	} else {
		for _, repository := range repositories.RemoteHrefs() {
			fmt.Fprintf(cobraCmd.OutOrStdout(), "%s\n", repository)
		}
		return nil
	}
}
