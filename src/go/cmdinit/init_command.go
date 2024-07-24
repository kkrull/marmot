package cmdinit

import (
	"fmt"

	"github.com/kkrull/marmot/cmd"
	"github.com/spf13/cobra"
)

// Construct a CLI command to initialize a meta repo
func NewInitCommand() *initCommand {
	return &initCommand{}
}

type initCommand struct{}

func (cliCmd *initCommand) RegisterWithCobra() {
	cobraCmd := cliCmd.toCobraCommand()
	cmd.AddMetaRepoCommand(*cobraCmd)
}

func (cliCmd *initCommand) toCobraCommand() *cobra.Command {
	return &cobra.Command{
		Args: cobra.NoArgs,
		Long: "Initialize a new Meta Repo, if none is already present.",
		RunE: func(cobraCmd *cobra.Command, _args []string) error {
			config := cmd.ParseFlags(cobraCmd)
			if config.Debug() {
				config.PrintDebug(cobraCmd.OutOrStdout())
				return nil
			}

			initAppCmd := config.AppFactory.InitCommand()
			metaRepoPath := config.MetaRepoPath()
			if runErr := initAppCmd.Run(metaRepoPath); runErr != nil {
				return runErr
			} else {
				fmt.Printf("Initialized meta repo at %s\n", metaRepoPath)
				return nil
			}
		},
		Short: "Initialize a meta repo",
		Use:   "init",
	}
}
