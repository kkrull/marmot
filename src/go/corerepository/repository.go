package corerepository

// Any number of Git repositories.
type Repositories interface {
	Names() []string
}

// One Git repository.
type Repository struct {
	Name string
}
