package userepository

import core "github.com/kkrull/marmot/corerepository"

// List repositories known to a meta repo.
type ListRemoteRepositoriesQuery = func() (core.Repositories, error)