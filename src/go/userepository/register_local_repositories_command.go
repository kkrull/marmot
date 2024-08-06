package userepository

import (
	"errors"

	core "github.com/kkrull/marmot/corerepository"
)

type RegisterLocalRepositoriesCommand struct {
	Source core.RepositorySource
}

func (cmd *RegisterLocalRepositoriesCommand) Run(localPaths ...string) error {
	return errors.ErrUnsupported
}
