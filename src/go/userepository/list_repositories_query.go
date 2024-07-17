package userepository

import core "github.com/kkrull/marmot/corerepository"

// List repositories known to a meta repo.
type ListRepositoriesQuery struct {
	Source core.RepositorySource
}

// TODO KDK: Rename this to remote repositories
func (cmd *ListRepositoriesQuery) Run() (core.Repositories, error) {
	return cmd.Source.List()
}
