package cmdinit

import (
	"fmt"

	"github.com/kkrull/marmot/cmd"
	"github.com/kkrull/marmot/usemetarepo"
	"github.com/spf13/cobra"
)

// Construct a CLI command to initialize a meta repo
func NewInitCommand(initApp *usemetarepo.InitCommand) *initCommand {
	return &initCommand{initAppCmd: initApp}
}

type initCommand struct {
	initAppCmd *usemetarepo.InitCommand
}

func (cliCmd *initCommand) RegisterWithCobra() {
	cmd.AddMetaRepoCommand(*cliCmd.toCobraCommand())
}

func (cliCmd *initCommand) toCobraCommand() *cobra.Command {
	return &cobra.Command{
		Long: "Initialize a new Meta Repo, if none is already present in the configured directory.",
		RunE: func(cobraCmd *cobra.Command, _args []string) error {
			config := cmd.ParseGlobalFlags(cobraCmd)
			fmt.Printf("- meta-home: %s\n", config.MetaRepoHome)

			if runErr := cliCmd.initAppCmd.Run(config.MetaRepoHome); runErr != nil {
				return fmt.Errorf("failed to initialize meta repo at %s; %w", config.MetaRepoHome, runErr)
			} else {
				fmt.Printf("Initialized meta repo at %s\n", config.MetaRepoHome)
				return nil
			}
		},
		Short: "Initialize a meta repo",
		Use:   "init",
	}
}
