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
	NewListLocalRepositories() (userepository.ListRemoteRepositoriesQuery, error)
	NewListRemoteRepositories() (userepository.ListRemoteRepositoriesQuery, error)
}

type queryFactory struct {
	MetaDataAdmin    coremetarepo.MetaDataAdmin
	RepositorySource corerepository.RepositorySource
}

func (factory *queryFactory) WithMetaDataAdmin(admin coremetarepo.MetaDataAdmin) *queryFactory {
	factory.MetaDataAdmin = admin
	return factory
}

func (factory *queryFactory) WithRepositorySource(source corerepository.RepositorySource) *queryFactory {
	factory.RepositorySource = source
	return factory
}

/* Repositories */

func (factory *queryFactory) NewListLocalRepositories() (userepository.ListRemoteRepositoriesQuery, error) {
	//Awaiting a way to register local repositories
	localRepositoriesAlwaysEmpty := func() (corerepository.Repositories, error) {
		repositories := make([]corerepository.Repository, 0)
		return &corerepository.RepositoriesArray{Repositories: repositories}, nil
	}

	return localRepositoriesAlwaysEmpty, nil
}

func (factory *queryFactory) NewListRemoteRepositories() (userepository.ListRemoteRepositoriesQuery, error) {
	if repositorySource, err := factory.repositorySource(); err != nil {
		return nil, err
	} else {
		return repositorySource.ListRemote, nil
	}
}

func (factory *queryFactory) repositorySource() (corerepository.RepositorySource, error) {
	if factory.RepositorySource == nil {
		return nil, errors.New("QueryFactory: missing RepositorySource")
	} else {
		return factory.RepositorySource, nil
	}
}
