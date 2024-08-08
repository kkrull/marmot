package userepository

import (
	"fmt"

	core "github.com/kkrull/marmot/corerepository"
)

// Registers Git repositories with the meta repo, based upon their paths on the local filesystem.
type RegisterLocalRepositoriesCommand struct {
	Source core.RepositorySource
}

func (cmd *RegisterLocalRepositoriesCommand) Run(localPaths ...string) error {
	if addErr := cmd.Source.AddLocals(localPaths); addErr != nil {
		return fmt.Errorf("failed to add local repositories: %w", addErr)
	} else {
		return nil
	}
}
