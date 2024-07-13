package core_repository

// List repositories known to a meta repo
type ListRepositoriesCommand struct {
	Source RepositorySource
}

func (cmd *ListRepositoriesCommand) Run() (Repositories, error) {
	return cmd.Source.List()
}
