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
	NewRegisterLocalRepositories() (*repository.RegisterLocalRepositoriesCommand, error)
	NewRegisterRemoteRepositories() (*repository.RegisterRemoteRepositoriesCommand, error)
}

type cmdFactory struct {
	LocalRepositorySource  corerepository.LocalRepositorySource
	MetaDataAdmin          coremetarepo.MetaDataAdmin
	RemoteRepositorySource corerepository.RemoteRepositorySource
}

func (factory *cmdFactory) WithLocalRepositorySource(source corerepository.LocalRepositorySource) *cmdFactory {
	factory.LocalRepositorySource = source
	return factory
}

func (factory *cmdFactory) WithMetaDataAdmin(admin coremetarepo.MetaDataAdmin) *cmdFactory {
	factory.MetaDataAdmin = admin
	return factory
}

func (factory *cmdFactory) WithRemoteRepositorySource(source corerepository.RemoteRepositorySource) *cmdFactory {
	factory.RemoteRepositorySource = source
	return factory
}

func (factory *cmdFactory) localRepositorySource() (corerepository.LocalRepositorySource, error) {
	if factory.LocalRepositorySource == nil {
		return nil, errors.New("missing LocalRepositorySource")
	} else {
		return factory.LocalRepositorySource, nil
	}
}

func (factory *cmdFactory) remoteRepositorySource() (corerepository.RemoteRepositorySource, error) {
	if factory.RemoteRepositorySource == nil {
		return nil, errors.New("missing RemoteRepositorySource")
	} else {
		return factory.RemoteRepositorySource, nil
	}
}

/* Administration */

func (factory *cmdFactory) NewInitMetaRepo() (*metarepo.InitCommand, error) {
	if factory.MetaDataAdmin == nil {
		return nil, errors.New("missing MetaDataAdmin")
	} else {
		return &metarepo.InitCommand{MetaDataAdmin: factory.MetaDataAdmin}, nil
	}
}

/* Repositories */

func (factory *cmdFactory) NewRegisterLocalRepositories() (
	*repository.RegisterLocalRepositoriesCommand, error,
) {
	if source, err := factory.localRepositorySource(); err != nil {
		return nil, err
	} else {
		return &repository.RegisterLocalRepositoriesCommand{Source: source}, nil
	}
}

func (factory *cmdFactory) NewRegisterRemoteRepositories() (
	*repository.RegisterRemoteRepositoriesCommand, error,
) {
	if source, err := factory.remoteRepositorySource(); err != nil {
		return nil, err
	} else {
		return &repository.RegisterRemoteRepositoriesCommand{Source: source}, nil
	}
}
