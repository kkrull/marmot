package userepository

import core "github.com/kkrull/marmot/corerepository"

// A source of Git repositories that a meta repo might care about
type RepositorySource interface {
	List() (core.Repositories, error)
}
