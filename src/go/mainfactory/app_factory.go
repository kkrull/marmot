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

func DefaultAppFactory() (*AppFactory, error) {
	if metaRepoPath, pathErr := defaultMetaRepoPath(); pathErr != nil {
		return nil, pathErr
	} else {
		return newAppFactory().ForLocalMetaRepo(metaRepoPath), nil
	}
}

func newAppFactory() *AppFactory {
	return &AppFactory{}
}

func defaultMetaRepoPath() (string, error) {
	if homeDir, homeErr := os.UserHomeDir(); homeErr != nil {
		return "", fmt.Errorf("failed to locate home directory; %w", homeErr)
	} else {
		return filepath.Join(homeDir, "meta"), nil
	}
}

// Constructs application commands and queries with configurable services.
type AppFactory struct {
	MetaDataAdmin    coremetarepo.MetaDataAdmin
	metaRepoPath     string
	RepositorySource corerepository.RepositorySource
}

// Configure a local, file-based meta repo at the specified path
func (factory *AppFactory) ForLocalMetaRepo(metaRepoPath string) *AppFactory {
	factory.metaRepoPath = metaRepoPath
	factory.RepositorySource = svcfs.NewJsonMetaRepo(metaRepoPath)
	return factory
}

func (factory *AppFactory) MetaRepoPath() string {
	return factory.metaRepoPath
}

/* Administration */

func (factory *AppFactory) InitCommand() (*metarepo.InitCommand, error) {
	if factory.MetaDataAdmin == nil {
		factory.MetaDataAdmin = svcfs.NewJsonMetaRepoAdmin()
	}

	return &metarepo.InitCommand{MetaDataAdmin: factory.MetaDataAdmin}, nil
}

/* Repositories */

func (factory *AppFactory) ListRemoteRepositoriesQuery() (repository.ListRemoteRepositoriesQuery, error) {
	if factory.RepositorySource == nil {
		return nil, errors.New("CommandFactory: missing RepositorySource")
	}

	return factory.RepositorySource.ListRemote, nil
}

func (factory *AppFactory) RegisterRemoteRepositoriesCommand() (
	*repository.RegisterRemoteRepositoriesCommand, error,
) {
	if factory.RepositorySource == nil {
		return nil, errors.New("CommandFactory: missing RepositorySource")
	}

	return &repository.RegisterRemoteRepositoriesCommand{Source: factory.RepositorySource}, nil
}
