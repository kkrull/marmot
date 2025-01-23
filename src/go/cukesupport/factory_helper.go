package cukesupport

import (
	"github.com/kkrull/marmot/svcfs"
	"github.com/kkrull/marmot/use"
)

// Create a factory for application commands on the current meta repo.
func ThatCommandFactory() (use.CommandFactory, error) {
	if metaRepoPath, pathErr := ThatMetaRepo(); pathErr != nil {
		return nil, pathErr
	} else {
		jsonMetaRepo := svcfs.NewJsonMetaRepo(metaRepoPath)
		return use.NewCommandFactory().WithRepositorySource(jsonMetaRepo), nil
	}
}

// Create a CommandFactory to initialize a meta repo with the specified version.
func ThatCommandFactoryS(version string) use.CommandFactory {
	return use.NewCommandFactory().
		WithMetaDataAdmin(svcfs.NewJsonMetaRepoAdmin(version))
}

// Create a factory to run queries on the current meta repo.
func ThatQueryFactory() (use.QueryFactory, error) {
	if metaRepoPath, pathErr := ThatMetaRepo(); pathErr != nil {
		return nil, pathErr
	} else {
		jsonMetaRepo := svcfs.NewJsonMetaRepo(metaRepoPath)
		return use.NewQueryFactory().
			WithLocalRepositorySource(jsonMetaRepo).
			WithRemoteRepositorySource(jsonMetaRepo), nil
	}
}
