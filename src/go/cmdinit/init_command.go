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
	cmd.AddMetaRepoCommand(*cliCmd.toCobraCommand())
}

func (cliCmd *initCommand) toCobraCommand() *cobra.Command {
	return &cobra.Command{
		Long: "Initialize a new Meta Repo, if none is already present in the configured directory.",
		RunE: func(cobraCmd *cobra.Command, _args []string) error {
			config := cmd.ParseFlags(cobraCmd)
			initAppCmd := config.AppFactory.InitCommand()
			if metaRepoPath, pathErr := config.MetaRepoPath(); pathErr != nil {
				return pathErr
			} else if runErr := initAppCmd.Run(metaRepoPath); runErr != nil {
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
