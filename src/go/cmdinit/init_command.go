package cmdinit

import (
	"fmt"
	"os"
	"path/filepath"

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
			if metaRepoPath, pathErr := defaultMetaRepoPath(); pathErr != nil {
				return pathErr
			} else if runErr := cliCmd.initAppCmd.Run(metaRepoPath); runErr != nil {
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

func defaultMetaRepoPath() (string, error) {
	if homeDir, homeErr := os.UserHomeDir(); homeErr != nil {
		return "", fmt.Errorf("failed to locate home directory; %w", homeErr)
	} else {
		return filepath.Join(homeDir, "meta"), nil
	}
}
