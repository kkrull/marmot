package mainfactory

import (
	"errors"

	"github.com/kkrull/marmot/coremetarepo"
	"github.com/kkrull/marmot/corerepository"
	"github.com/kkrull/marmot/svcfs"
	metarepo "github.com/kkrull/marmot/usemetarepo"
	repository "github.com/kkrull/marmot/userepository"
)

// Constructs application commands and queries with configurable services.
type AppFactory struct {
	MetaDataAdmin    coremetarepo.MetaDataAdmin
	RepositorySource corerepository.RepositorySource
}

// Configure a local, file-based meta repo at the specified path
func (factory *AppFactory) ForLocalMetaRepo(metaRepoPath string) *AppFactory {
	factory.RepositorySource = svcfs.NewJsonMetaRepo(metaRepoPath)
	return factory
}

/* Administration */

func (factory *AppFactory) InitCommand() (*metarepo.InitCommand, error) {
	if factory.MetaDataAdmin == nil {
		factory.MetaDataAdmin = svcfs.NewJsonMetaRepoAdmin()
	}

	return &metarepo.InitCommand{MetaDataAdmin: factory.MetaDataAdmin}, nil
}

/* Repositories */

func (factory *AppFactory) ListRemoteRepositoriesQuery() (repository.ListRemoteRepositoriesQuery, error) {
	if factory.RepositorySource == nil {
		return nil, errors.New("CommandFactory: missing RepositorySource")
	}

	return factory.RepositorySource.ListRemote, nil
}

func (factory *AppFactory) RegisterRemoteRepositoriesCommand() (
	*repository.RegisterRemoteRepositoriesCommand, error,
) {
	if factory.RepositorySource == nil {
		return nil, errors.New("CommandFactory: missing RepositorySource")
	}

	return &repository.RegisterRemoteRepositoriesCommand{Source: factory.RepositorySource}, nil
}
