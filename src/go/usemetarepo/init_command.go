package usemetarepo

import core "github.com/kkrull/marmot/coremetarepo"

// Initializes a new meta repo where none existed before.
type InitCommand struct {
	MetaDataAdmin core.MetaDataAdmin
}

func (cmd InitCommand) Run(metaRepoPath string) error {
	return cmd.MetaDataAdmin.Init(metaRepoPath)
}
