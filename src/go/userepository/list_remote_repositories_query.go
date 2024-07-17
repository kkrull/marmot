package userepository

import core "github.com/kkrull/marmot/corerepository"

// List repositories known to a meta repo.
type ListRemoteRepositoriesQuery struct {
	Source core.RepositorySource
}

func (cmd *ListRemoteRepositoriesQuery) Run() (core.Repositories, error) {
	return cmd.Source.List()
}
