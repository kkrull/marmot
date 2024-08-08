package userepository

import (
	"fmt"
	"path/filepath"

	core "github.com/kkrull/marmot/corerepository"
)

// Registers Git repositories with the meta repo, based upon their paths on the local filesystem.
type RegisterLocalRepositoriesCommand struct {
	Source core.RepositorySource
}

func (cmd *RegisterLocalRepositoriesCommand) Run(localPaths []string) error {
	if addErr := cmd.Source.AddLocals(normalizePaths(localPaths)); addErr != nil {
		return fmt.Errorf("failed to add local repositories; %w", addErr)
	} else {
		return nil
	}
}

func normalizePaths(rawPaths []string) []string {
	normalized := make([]string, len(rawPaths))
	for i, rawPath := range rawPaths {
		if normalPath, pathErr := filepath.Abs(rawPath); pathErr == nil {
			normalized[i] = normalPath
		}
	}

	return normalized
}
