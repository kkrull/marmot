package cmd

import (
	cmdlocal "github.com/kkrull/marmot/cmd/local"
	cmdroot "github.com/kkrull/marmot/cmd/root"
	"github.com/spf13/cobra"
)

// Construct a CLI command to deal with repositories on the local filesystem.
func NewLocalCommand() *localCommand {
	return &localCommand{}
}

type localCommand struct{}

/* Mapping to Cobra */

// Add this command as a sub-command of the given Cobra command.
func (cliCmd *localCommand) AddToCobra(cobraCmd *cobra.Command) {
	cobraCmd.AddCommand(cliCmd.toCobraCommand())
}

func (cliCmd *localCommand) toCobraCommand() *cobra.Command {
	localCobraCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: repositoryGroup.id(),
		Long:    "Deal with repositories on the local filesystem.",
		RunE:    runLocalCobra,
		Short:   "Deal with local repositories",
		Use:     "local",
	}

	// TODO KDK: Use a separate .marmot/meta-repo-local.json file
	cmdlocal.NewListCommand().AddToCobra(localCobraCmd)
	cmdlocal.NewRegisterCommand().AddToCobra(localCobraCmd)
	return localCobraCmd
}

func runLocalCobra(cli *cobra.Command, args []string) error {
	if parser, newErr := cmdroot.RootConfigParser(); newErr != nil {
		return newErr
	} else if config, parseErr := parser.Parse(cli.Flags(), args); parseErr != nil {
		return parseErr
	} else if config.Debug() {
		config.PrintDebug(cli.OutOrStdout())
		return nil
	} else if len(args) == 0 {
		return cli.Help()
	} else {
		// Run the sub-command named in the arguments
		return nil
	}
}
