package userepository

import (
	"fmt"
	"path/filepath"
	"slices"

	core "github.com/kkrull/marmot/corerepository"
)

// Registers Git repositories with the meta repo, based upon their paths on the local filesystem.
type RegisterLocalRepositoriesCommand struct {
	Source core.RepositorySource
}

func (cmd *RegisterLocalRepositoriesCommand) Run(localPaths []string) error {
	distinctNormalizedPaths := make([]string, 0)
	for _, rawPath := range localPaths {
		if absPath, absErr := filepath.Abs(rawPath); absErr != nil {
			return absErr
		} else if isDuplicate := slices.Contains(distinctNormalizedPaths, absPath); isDuplicate {
			continue
		} else {
			distinctNormalizedPaths = append(distinctNormalizedPaths, absPath)
		}
	}

	if addErr := cmd.Source.AddLocals(distinctNormalizedPaths); addErr != nil {
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
