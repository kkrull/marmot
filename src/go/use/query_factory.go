package use

import (
	"errors"

	"github.com/kkrull/marmot/coremetarepo"
	"github.com/kkrull/marmot/corerepository"
	"github.com/kkrull/marmot/userepository"
)

func NewQueryFactory() *queryFactory {
	return &queryFactory{}
}

// Constructs application queries with configurable services.
type QueryFactory interface {
	NewListLocalRepositories() (userepository.ListLocalRepositoriesQuery, error)
	NewListRemoteRepositories() (userepository.ListRemoteRepositoriesQuery, error)
}

type queryFactory struct {
	LocalRepositorySource  corerepository.LocalRepositorySource
	MetaDataAdmin          coremetarepo.MetaDataAdmin
	RemoteRepositorySource corerepository.RemoteRepositorySource
}

func (factory *queryFactory) WithLocalRepositorySource(source corerepository.LocalRepositorySource) *queryFactory {
	factory.LocalRepositorySource = source
	return factory
}

func (factory *queryFactory) WithMetaDataAdmin(admin coremetarepo.MetaDataAdmin) *queryFactory {
	factory.MetaDataAdmin = admin
	return factory
}

func (factory *queryFactory) WithRemoteRepositorySource(source corerepository.RemoteRepositorySource) *queryFactory {
	factory.RemoteRepositorySource = source
	return factory
}

/* Repositories */

func (factory *queryFactory) NewListLocalRepositories() (userepository.ListLocalRepositoriesQuery, error) {
	if repositorySource, err := factory.localRepositorySource(); err != nil {
		return nil, err
	} else {
		return repositorySource.ListLocal, nil
	}
}

func (factory *queryFactory) NewListRemoteRepositories() (userepository.ListRemoteRepositoriesQuery, error) {
	if repositorySource, err := factory.remoteRepositorySource(); err != nil {
		return nil, err
	} else {
		return repositorySource.ListRemote, nil
	}
}

func (factory *queryFactory) localRepositorySource() (corerepository.LocalRepositorySource, error) {
	if factory.LocalRepositorySource == nil {
		return nil, errors.New("missing LocalRepositorySource")
	} else {
		return factory.LocalRepositorySource, nil
	}
}

func (factory *queryFactory) remoteRepositorySource() (corerepository.RemoteRepositorySource, error) {
	if factory.RemoteRepositorySource == nil {
		return nil, errors.New("missing RemoteRepositorySource")
	} else {
		return factory.RemoteRepositorySource, nil
	}
}
