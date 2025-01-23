package userepository

import (
	"fmt"
	"path/filepath"

	core "github.com/kkrull/marmot/corerepository"
)

// Registers Git repositories with the meta repo, based upon their paths on the local filesystem.
type RegisterLocalRepositoriesCommand struct {
	Source core.LocalRepositorySource
}

func (cmd *RegisterLocalRepositoriesCommand) Run(localPaths []string) error {
	absolutePaths := make([]string, len(localPaths))
	for i, rawPath := range localPaths {
		if absPath, absErr := filepath.Abs(rawPath); absErr != nil {
			return absErr
		} else {
			absolutePaths[i] = absPath
		}
	}

	if addErr := cmd.Source.AddLocals(absolutePaths); addErr != nil {
		return fmt.Errorf("failed to add local repositories; %w", addErr)
	} else {
		return nil
	}
}
