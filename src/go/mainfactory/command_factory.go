package mainfactory

import (
	"errors"
	"net"

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

// Configure a local, file-based meta repo at the specified path
func (factory *CommandFactory) WithJsonFileSource(metaRepoPath string) {
	metaRepo := svcfs.NewJsonMetaDataRepo(metaRepoPath)
	factory.MetaDataAdmin = metaRepo
	factory.RepositorySource = metaRepo
}

/* Meta repo administration */

func (factory *CommandFactory) InitCommand() (*usemetarepo.InitCommand, error) {
	if factory.MetaDataAdmin == nil {
		return nil, errors.New("CommandFactory: missing MetaDataAdmin")
	}

	return &usemetarepo.InitCommand{
		MetaDataAdmin: factory.MetaDataAdmin,
	}, nil
}

/* Repositories */

func (factory *CommandFactory) ListRepositoriesQuery() (*userepository.ListRepositoriesQuery, error) {
	if factory.RepositorySource == nil {
		return nil, errors.New("CommandFactory: missing RepositorySource")
	}

	return &userepository.ListRepositoriesQuery{
		Source: factory.RepositorySource,
	}, nil
}

func (factory *CommandFactory) RegisterRemoteRepositoriesCommand(hostUrls []net.Addr) (interface{}, error) {
	return 42, nil
}
