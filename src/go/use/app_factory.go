package use

import (
	"errors"

	"github.com/kkrull/marmot/coremetarepo"
	"github.com/kkrull/marmot/corerepository"
	metarepo "github.com/kkrull/marmot/usemetarepo"
	repository "github.com/kkrull/marmot/userepository"
)

func NewAppFactory() *appFactory {
	return &appFactory{}
}

// Constructs application commands and queries with configurable services.
type AppFactory interface {
	InitCommand() (*metarepo.InitCommand, error)
	ListRemoteRepositoriesQuery() (repository.ListRemoteRepositoriesQuery, error)
	RegisterRemoteRepositoriesCommand() (*repository.RegisterRemoteRepositoriesCommand, error)
}

type appFactory struct {
	MetaDataAdmin    coremetarepo.MetaDataAdmin
	RepositorySource corerepository.RepositorySource
}

func (factory *appFactory) WithMetaDataAdmin(metadataAdmin coremetarepo.MetaDataAdmin) *appFactory {
	factory.MetaDataAdmin = metadataAdmin
	return factory
}

func (factory *appFactory) WithRepositorySource(repositorySource corerepository.RepositorySource) *appFactory {
	factory.RepositorySource = repositorySource
	return factory
}

/* Administration */

func (factory *appFactory) InitCommand() (*metarepo.InitCommand, error) {
	if factory.MetaDataAdmin == nil {
		return nil, errors.New("AppFactory: missing MetaDataAdmin")
	} else {
		return &metarepo.InitCommand{MetaDataAdmin: factory.MetaDataAdmin}, nil
	}
}

/* Repositories */

func (factory *appFactory) ListRemoteRepositoriesQuery() (repository.ListRemoteRepositoriesQuery, error) {
	if repositorySource, err := factory.repositorySource(); err != nil {
		return nil, err
	} else {
		return repositorySource.ListRemote, nil
	}
}

func (factory *appFactory) RegisterRemoteRepositoriesCommand() (
	*repository.RegisterRemoteRepositoriesCommand, error,
) {
	if repositorySource, err := factory.repositorySource(); err != nil {
		return nil, err
	} else {
		return &repository.RegisterRemoteRepositoriesCommand{Source: repositorySource}, nil
	}
}

func (factory *appFactory) repositorySource() (corerepository.RepositorySource, error) {
	if factory.RepositorySource == nil {
		return nil, errors.New("AppFactory: missing RepositorySource")
	} else {
		return factory.RepositorySource, nil
	}
}
