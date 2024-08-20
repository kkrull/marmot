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
	absolutePaths := make([]string, len(localPaths))
	for i, rawPath := range localPaths {
		// if _, statErr := os.Stat(rawPath); os.IsNotExist(statErr) {
		// 	return fmt.Errorf("path does not exist")
		// }

		if absPath, absErr := filepath.Abs(rawPath); absErr != nil {
			//Happens when os.Getwd fails, such as if the current working directory no longer exists.
			//However, this is difficult to test in a repeatable, platform-independent way.
			//https://stackoverflow.com/a/75753434/112682
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
