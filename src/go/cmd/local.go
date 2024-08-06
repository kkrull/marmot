package cmd

import (
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
	remoteCobraCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		GroupID: repositoryGroup.id(),
		Long:    "Deal with repositories on the local filesystem.",
		RunE:    runLocalCobra,
		Short:   "Deal with local repositories",
		Use:     "local",
	}

	// cmdremote.NewListCommand().AddToCobra(remoteCobraCmd)
	// cmdremote.NewRegisterCommand().AddToCobra(remoteCobraCmd)
	return remoteCobraCmd
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
