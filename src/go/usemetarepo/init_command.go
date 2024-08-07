package usemetarepo

import (
	"fmt"

	core "github.com/kkrull/marmot/coremetarepo"
)

// Initializes a new meta repo where none existed before.
type InitCommand struct {
	MetaDataAdmin core.MetaDataAdmin
}

func (cmd InitCommand) Run(metaRepoPath string) error {
	if cmd.MetaDataAdmin.Exists(metaRepoPath) {
		return fmt.Errorf("%s is already a meta repo", metaRepoPath)
	} else if createErr := cmd.MetaDataAdmin.Create(metaRepoPath); createErr != nil {
		return fmt.Errorf("failed to initialize meta repo; %w", createErr)
	} else {
		return nil
	}
}
