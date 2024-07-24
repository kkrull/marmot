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
			metaHomeFlag := cobraCmd.Flags().Lookup("meta-home")
			metaHomePath := metaHomeFlag.Value.String()
			fmt.Printf("[init_command] meta-home: %s\n", metaHomePath)

			if runErr := cliCmd.initAppCmd.Run(metaHomePath); runErr != nil {
				return fmt.Errorf("failed to initialize meta repo at %s; %w", metaHomePath, runErr)
			} else {
				fmt.Printf("Initialized meta repo at %s\n", metaHomePath)
				return nil
			}
		},
		Short: "Initialize a meta repo",
		Use:   "init",
	}
}
