package corerepository

// A source of Git repositories that a meta repo might care about.
type RepositorySource interface {
	List() (Repositories, error)
}
