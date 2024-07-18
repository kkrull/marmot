package mainfactory

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kkrull/marmot/coremetarepo"
	"github.com/kkrull/marmot/corerepository"
	"github.com/kkrull/marmot/svcfs"
	metarepo "github.com/kkrull/marmot/usemetarepo"
	repository "github.com/kkrull/marmot/userepository"
)

func DefaultCommandQueryFactory() (*CommandQueryFactory, error) {
	if metaRepoPath, pathErr := DefaultMetaRepoPath(); pathErr != nil {
		return nil, pathErr
	} else {
		factory := &CommandQueryFactory{}
		factory.ForLocalMetaRepo(metaRepoPath)
		return factory, nil
	}
}

func DefaultMetaRepoPath() (string, error) {
	if homeDir, homeErr := os.UserHomeDir(); homeErr != nil {
		return "", fmt.Errorf("failed to locate home directory; %w", homeErr)
	} else {
		return filepath.Join(homeDir, "meta"), nil
	}
}

// Constructs commands and queries with configurable dependencies.
type CommandQueryFactory struct {
	MetaDataAdmin    coremetarepo.MetaDataAdmin
	RepositorySource corerepository.RepositorySource
}

// Configure a local, file-based meta repo at the specified path
func (factory *CommandQueryFactory) ForLocalMetaRepo(metaRepoPath string) {
	factory.RepositorySource = svcfs.NewJsonMetaRepo(metaRepoPath)
}

/* Administration */

func (factory *CommandQueryFactory) InitCommand() (*metarepo.InitCommand, error) {
	if factory.MetaDataAdmin == nil {
		factory.MetaDataAdmin = svcfs.NewJsonMetaRepoAdmin()
	}

	return &metarepo.InitCommand{MetaDataAdmin: factory.MetaDataAdmin}, nil
}

/* Repositories */

func (factory *CommandQueryFactory) ListRemoteRepositoriesQuery() (repository.ListRemoteRepositoriesQuery, error) {
	if factory.RepositorySource == nil {
		return nil, errors.New("CommandFactory: missing RepositorySource")
	}

	return factory.RepositorySource.ListRemote, nil
}

func (factory *CommandQueryFactory) RegisterRemoteRepositoriesCommand() (
	*repository.RegisterRemoteRepositoriesCommand, error,
) {
	if factory.RepositorySource == nil {
		return nil, errors.New("CommandFactory: missing RepositorySource")
	}

	return &repository.RegisterRemoteRepositoriesCommand{Source: factory.RepositorySource}, nil
}
