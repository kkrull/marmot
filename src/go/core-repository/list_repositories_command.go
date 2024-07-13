package core_repository

import "errors"

// List repositories known to a meta repo
type ListRepositoriesCommand struct {}

func (cmd *ListRepositoriesCommand) Run() ([]string, error) {
	return nil, errors.New("not implemented")
}
