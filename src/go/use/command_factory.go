package use

import (
	"errors"

	"github.com/kkrull/marmot/coremetarepo"
	"github.com/kkrull/marmot/corerepository"
	metarepo "github.com/kkrull/marmot/usemetarepo"
	repository "github.com/kkrull/marmot/userepository"
)

func NewActionFactory() *actionFactory {
	return &actionFactory{}
}

// Constructs application actions and queries with configurable services.
type ActionFactory interface {
	NewInitMetaRepo() (*metarepo.InitAction, error)
	NewRegisterLocalRepositories() (*repository.RegisterLocalRepositoriesAction, error)
	NewRegisterRemoteRepositories() (*repository.RegisterRemoteRepositoriesAction, error)
}

type actionFactory struct {
	LocalRepositorySource  corerepository.LocalRepositorySource
	MetaDataAdmin          coremetarepo.MetaDataAdmin
	RemoteRepositorySource corerepository.RemoteRepositorySource
}

func (factory *actionFactory) WithLocalRepositorySource(source corerepository.LocalRepositorySource) *actionFactory {
	factory.LocalRepositorySource = source
	return factory
}

func (factory *actionFactory) WithMetaDataAdmin(admin coremetarepo.MetaDataAdmin) *actionFactory {
	factory.MetaDataAdmin = admin
	return factory
}

func (factory *actionFactory) WithRemoteRepositorySource(source corerepository.RemoteRepositorySource) *actionFactory {
	factory.RemoteRepositorySource = source
	return factory
}

func (factory *actionFactory) localRepositorySource() (corerepository.LocalRepositorySource, error) {
	if factory.LocalRepositorySource == nil {
		return nil, errors.New("missing LocalRepositorySource")
	} else {
		return factory.LocalRepositorySource, nil
	}
}

func (factory *actionFactory) remoteRepositorySource() (corerepository.RemoteRepositorySource, error) {
	if factory.RemoteRepositorySource == nil {
		return nil, errors.New("missing RemoteRepositorySource")
	} else {
		return factory.RemoteRepositorySource, nil
	}
}

/* Administration */

func (factory *actionFactory) NewInitMetaRepo() (*metarepo.InitAction, error) {
	if factory.MetaDataAdmin == nil {
		return nil, errors.New("missing MetaDataAdmin")
	} else {
		return &metarepo.InitAction{MetaDataAdmin: factory.MetaDataAdmin}, nil
	}
}

/* Repositories */

func (factory *actionFactory) NewRegisterLocalRepositories() (
	*repository.RegisterLocalRepositoriesAction, error,
) {
	if source, err := factory.localRepositorySource(); err != nil {
		return nil, err
	} else {
		return &repository.RegisterLocalRepositoriesAction{Source: source}, nil
	}
}

func (factory *actionFactory) NewRegisterRemoteRepositories() (
	*repository.RegisterRemoteRepositoriesAction, error,
) {
	if source, err := factory.remoteRepositorySource(); err != nil {
		return nil, err
	} else {
		return &repository.RegisterRemoteRepositoriesAction{Source: source}, nil
	}
}
