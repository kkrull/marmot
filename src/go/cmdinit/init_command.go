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
			if config, parseErr := cmd.ParseFlags(cobraCmd); parseErr != nil {
				return parseErr
			} else if config.Debug {
				config.PrintDebug(cobraCmd.OutOrStdout())
				return nil
			} else {
				initAppCmd := config.AppFactory.InitCommand()
				if runErr := initAppCmd.Run(config.MetaRepoPath); runErr != nil {
					return runErr
				} else {
					fmt.Printf("Initialized meta repo at %s\n", config.MetaRepoPath)
					return nil
				}
			}
		},
		Short: "Initialize a meta repo",
		Use:   "init",
	}
}
