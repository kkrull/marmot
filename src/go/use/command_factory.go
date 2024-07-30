package use

import (
	"errors"

	"github.com/kkrull/marmot/coremetarepo"
	"github.com/kkrull/marmot/corerepository"
	metarepo "github.com/kkrull/marmot/usemetarepo"
	repository "github.com/kkrull/marmot/userepository"
)

func NewCommandFactory() *cmdFactory {
	return &cmdFactory{}
}

// Constructs application commands and queries with configurable services.
type CommandFactory interface {
	NewInitMetaRepo() (*metarepo.InitCommand, error)
	NewRegisterRemoteRepositories() (*repository.RegisterRemoteRepositoriesCommand, error)
}

type cmdFactory struct {
	MetaDataAdmin    coremetarepo.MetaDataAdmin
	RepositorySource corerepository.RepositorySource
}

func (factory *cmdFactory) WithMetaDataAdmin(metadataAdmin coremetarepo.MetaDataAdmin) *cmdFactory {
	factory.MetaDataAdmin = metadataAdmin
	return factory
}

func (factory *cmdFactory) WithRepositorySource(repositorySource corerepository.RepositorySource) *cmdFactory {
	factory.RepositorySource = repositorySource
	return factory
}

func (factory *cmdFactory) repositorySource() (corerepository.RepositorySource, error) {
	if factory.RepositorySource == nil {
		return nil, errors.New("AppFactory: missing RepositorySource")
	} else {
		return factory.RepositorySource, nil
	}
}

/* Administration */

func (factory *cmdFactory) NewInitMetaRepo() (*metarepo.InitCommand, error) {
	if factory.MetaDataAdmin == nil {
		return nil, errors.New("AppFactory: missing MetaDataAdmin")
	} else {
		return &metarepo.InitCommand{MetaDataAdmin: factory.MetaDataAdmin}, nil
	}
}

/* Remote repositories */

func (factory *cmdFactory) NewRegisterRemoteRepositories() (
	*repository.RegisterRemoteRepositoriesCommand, error,
) {
	if repositorySource, err := factory.repositorySource(); err != nil {
		return nil, err
	} else {
		return &repository.RegisterRemoteRepositoriesCommand{Source: repositorySource}, nil
	}
}
