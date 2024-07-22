package cmd

import (
	"fmt"

	"github.com/kkrull/marmot/mainfactory"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	GroupID: metaRepoGroup,
	Long:    "Initialize a blank Meta Repo in the configured directory, if none is already present.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if factory, factoryErr := mainfactory.DefaultCommandQueryFactory(); factoryErr != nil {
			return factoryErr
		} else if initUseCmd, initErr := factory.InitCommand(); initErr != nil {
			return fmt.Errorf("failed to create command; %w", initErr)
		} else if metaRepoHome, pathErr := mainfactory.DefaultMetaRepoPath(); pathErr != nil {
			return pathErr
		} else if runErr := initUseCmd.Run(metaRepoHome); runErr != nil {
			return fmt.Errorf("failed to initialize meta repo; %w", runErr)
		} else {
			fmt.Printf("Initialized meta repo at %s\n", metaRepoHome)
			return nil
		}
	},
	Short: "initialize a meta repo",
	Use:   "init",
}

func init() {
	addCommandToRoot(initCmd)
}
