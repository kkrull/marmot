package main_factory

import (
	"errors"

	metarepo "github.com/kkrull/marmot/core-metarepo"
	"github.com/kkrull/marmot/fs"
)

// Constructs commands with configurable dependencies
type CommandFactory struct {
	MetaDataSource metarepo.MetaDataSource
}

// Construct a command to initialize a new meta repo
func (factory *CommandFactory) NewInitCommand() (*metarepo.InitCommand, error) {
	if factory.MetaDataSource == nil {
		return nil, errors.New("CommandFactory: missing MetaDataSource")
	}

	return nil, nil
}

func (factory *CommandFactory) WithJsonFileSource(metaRepoPath string) {
	factory.MetaDataSource = fs.JsonMetaDataSource(metaRepoPath)
}
