package mainfactory

import (
	"errors"

	"github.com/cucumber/godog"
	"github.com/kkrull/marmot/coremetarepo"
	"github.com/kkrull/marmot/corerepository"
	"github.com/kkrull/marmot/svcfs"
	"github.com/kkrull/marmot/usemetarepo"
	"github.com/kkrull/marmot/userepository"
)

// Constructs commands with configurable dependencies.
type CommandFactory struct {
	MetaDataAdmin    coremetarepo.MetaDataAdmin
	RepositorySource corerepository.RepositorySource
}

func (factory *CommandFactory) InitCommand() (*usemetarepo.InitCommand, error) {
	if factory.MetaDataAdmin == nil {
		return nil, errors.New("CommandFactory: missing MetaDataAdmin")
	}

	return &usemetarepo.InitCommand{
		MetaDataAdmin: factory.MetaDataAdmin,
	}, nil
}

func (factory *CommandFactory) ListRepositoriesQuery() (*userepository.ListRepositoriesQuery, error) {
	if factory.RepositorySource == nil {
		return nil, godog.ErrPending
		// return nil, errors.New("CommandFactory: missing RepositorySource")
	}

	return &userepository.ListRepositoriesQuery{
		Source: factory.RepositorySource,
	}, nil
}

func (factory *CommandFactory) WithJsonFileSource(metaRepoPath string) {
	factory.MetaDataAdmin = svcfs.JsonMetaDataRepo(metaRepoPath)
}
