package userepository

import core "github.com/kkrull/marmot/corerepository"

// List registered repositories that have been cloned to the local file system.
type ListLocalRepositoriesQuery = func() (core.Repositories, error)

// List registered repositories on remote hosts.
type ListRemoteRepositoriesQuery = func() (core.Repositories, error)
