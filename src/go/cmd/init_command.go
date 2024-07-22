package cmd

import (
	"fmt"

	"github.com/kkrull/marmot/mainfactory" //TODO KDK: Stop depending upon this
	"github.com/spf13/cobra"
)

// TODO KDK: Move to package that depends upon cmd (root)
var initCmd = &cobra.Command{
	GroupID: metaRepoGroup,
	Long:    "Initialize a blank Meta Repo in the configured directory, if none is already present.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if appFactory, configErr := mainfactory.DefaultAppFactory(); configErr != nil {
			return configErr
		} else if initAppCmd, buildErr := appFactory.InitCommand(); buildErr != nil {
			return fmt.Errorf("failed to create command; %w", buildErr)
		} else if metaRepoHome, pathErr := mainfactory.DefaultMetaRepoPath(); pathErr != nil {
			return pathErr
		} else if runErr := initAppCmd.Run(metaRepoHome); runErr != nil {
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
	AddMetaRepoCommand(initCmd)
}
