package userepository

import core "github.com/kkrull/marmot/corerepository"

// List repositories known to a meta repo.
type ListRepositoriesCommand struct {
	Source core.RepositorySource
}

func (cmd *ListRepositoriesCommand) Run() (core.Repositories, error) {
	return cmd.Source.List()
}
