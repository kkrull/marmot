package main_factory

import (
	"errors"

	metarepo "github.com/kkrull/marmot/core-metarepo"
	repository "github.com/kkrull/marmot/core-repository"
	"github.com/kkrull/marmot/fs"
)

// Constructs commands with configurable dependencies
type CommandFactory struct {
	MetaDataSource   metarepo.MetaDataSource
	RepositorySource repository.RepositorySource
}

func (factory *CommandFactory) InitCommand() (*metarepo.InitCommand, error) {
	if factory.MetaDataSource == nil {
		return nil, errors.New("CommandFactory: missing MetaDataSource")
	}

	return &metarepo.InitCommand{
		MetaDataSource: factory.MetaDataSource,
	}, nil
}

func (factory *CommandFactory) ListRepositoriesCommand() (*repository.ListRepositoriesCommand, error) {
	if factory.RepositorySource == nil {
		return nil, errors.New("CommandFactory: missing RepositorySource")
	}

	return &repository.ListRepositoriesCommand{
		Source: factory.RepositorySource,
	}, nil
}

func (factory *CommandFactory) WithJsonFileSource(metaRepoPath string) {
	factory.MetaDataSource = fs.JsonMetaDataSource(metaRepoPath)
}
