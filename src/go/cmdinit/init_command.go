package cmdinit

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kkrull/marmot/cmd"
	"github.com/spf13/cobra"
)

// Construct a CLI command to initialize a meta repo
func NewInitCommand() *initCommand {
	return &initCommand{}
}

type initCommand struct{}

func (cliCmd *initCommand) RegisterWithCobra() error {
	if cobraCmd, cobraErr := cliCmd.toCobraCommand(); cobraErr != nil {
		return cobraErr
	} else {
		cmd.AddMetaRepoCommand(*cobraCmd)
		return nil
	}
}

func (cliCmd *initCommand) toCobraCommand() (*cobra.Command, error) {
	cobraCmd := &cobra.Command{
		Long: "Initialize a new Meta Repo, if none is already present.",
		RunE: func(cobraCmd *cobra.Command, _args []string) error {
			config := cmd.ParseFlags(cobraCmd)
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

	if defaultPath, pathErr := defaultMetaRepoPath(); pathErr != nil {
		return nil, pathErr
	} else {
		cobraCmd.Flags().String("meta-repo", defaultPath, "Meta repo to use")
		return cobraCmd, nil
	}
}

func defaultMetaRepoPath() (string, error) {
	if homeDir, homeErr := os.UserHomeDir(); homeErr != nil {
		return "", fmt.Errorf("failed to locate home directory; %w", homeErr)
	} else {
		return filepath.Join(homeDir, "meta"), nil
	}
}
