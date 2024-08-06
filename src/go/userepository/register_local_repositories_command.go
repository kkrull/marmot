package userepository

import (
	core "github.com/kkrull/marmot/corerepository"
)

// Registers Git repositories with the meta repo, based upon their paths on the local filesystem.
type RegisterLocalRepositoriesCommand struct {
	Source core.RepositorySource
}

func (cmd *RegisterLocalRepositoriesCommand) Run(localPaths ...string) error {
	for _, localPath := range localPaths {
		cmd.Source.AddLocal(localPath)
	}

	return nil
}
