package mainfactory

import (
	"errors"

	"github.com/kkrull/marmot/coremetarepo"
	"github.com/kkrull/marmot/corerepository"
	"github.com/kkrull/marmot/svcfs"
	metarepo "github.com/kkrull/marmot/usemetarepo"
	repository "github.com/kkrull/marmot/userepository"
)

// Constructs commands with configurable dependencies.
type CommandFactory struct {
	MetaDataAdmin    coremetarepo.MetaDataAdmin
	RepositorySource corerepository.RepositorySource
}

// Configure a file-based meta repo at the specified path on the local file system
// func (factory *CommandFactory) ForLocalPath(metaRepoPath string) {
// 	// metaRepo := svcfs.NewJsonMetaDataRepo(metaRepoPath)
// 	// factory.MetaDataAdmin = metaRepo
// 	// factory.RepositorySource = metaRepo
// }

// Configure a local, file-based meta repo at the specified path
func (factory *CommandFactory) WithJsonFileSource(metaRepoPath string) {
	metaRepo := svcfs.NewJsonMetaDataRepo(metaRepoPath)
	factory.MetaDataAdmin = metaRepo
	factory.RepositorySource = metaRepo
}

/* Meta repo administration */

func (factory *CommandFactory) InitCommand() (*metarepo.InitCommand, error) {
	if factory.MetaDataAdmin == nil {
		return nil, errors.New("CommandFactory: missing MetaDataAdmin")
	}

	return &metarepo.InitCommand{
		MetaDataAdmin: factory.MetaDataAdmin,
	}, nil
}

func (factory *CommandFactory) InitCommandP(metaRepoPath string) (*metarepo.InitCommand, error) {
	if factory.MetaDataAdmin == nil {
		return nil, errors.New("CommandFactory: missing MetaDataAdmin")
	}

	return &metarepo.InitCommand{
		MetaDataAdmin: factory.MetaDataAdmin,
	}, nil
}

/* Repositories */

func (factory *CommandFactory) ListRepositoriesQuery() (*repository.ListRemoteRepositoriesQuery, error) {
	if factory.RepositorySource == nil {
		return nil, errors.New("CommandFactory: missing RepositorySource")
	}

	return &repository.ListRemoteRepositoriesQuery{
		Source: factory.RepositorySource,
	}, nil
}

func (factory *CommandFactory) RegisterRemoteRepositoriesCommand() (
	*repository.RegisterRemoteRepositoriesCommand, error,
) {
	if factory.RepositorySource == nil {
		return nil, errors.New("CommandFactory: missing RepositorySource")
	}

	return &repository.RegisterRemoteRepositoriesCommand{
		Source: factory.RepositorySource,
	}, nil
}
