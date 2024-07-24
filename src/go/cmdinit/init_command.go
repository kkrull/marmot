package cmdinit

import (
	"fmt"

	"github.com/kkrull/marmot/cmd"
	"github.com/kkrull/marmot/usemetarepo"
	"github.com/spf13/cobra"
)

// Construct a CLI command to initialize a meta repo at the specified path
func NewInitCommand(initApp *usemetarepo.InitCommand, metaRepoHome string) *initCommand {
	return &initCommand{
		initApp: initApp,
		path:    metaRepoHome,
	}
}

type initCommand struct {
	initApp *usemetarepo.InitCommand
	path    string
}

func (cliCmd *initCommand) RegisterWithCobra() {
	cmd.AddMetaRepoCommand(*cliCmd.toCobraCommand())
}

func (cliCmd *initCommand) toCobraCommand() *cobra.Command {
	return &cobra.Command{
		Long: "Initialize a new Meta Repo in the configured directory, if none is already present.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if runErr := cliCmd.initApp.Run(cliCmd.path); runErr != nil {
				return fmt.Errorf("failed to initialize meta repo at %s; %w", cliCmd.path, runErr)
			} else {
				fmt.Printf("Initialized meta repo at %s\n", cliCmd.path)
				return nil
			}
		},
		Short: "Initialize a meta repo",
		Use:   "init",
	}
}
