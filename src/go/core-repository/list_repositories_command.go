package core_repository

import (
	"github.com/cucumber/godog"
)

// List repositories known to a meta repo
type ListRepositoriesCommand struct {}

func (cmd *ListRepositoriesCommand) Run() (Repositories, error) {
	return nil, godog.ErrPending
}
