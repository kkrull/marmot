package userepository

import (
	"fmt"
	"net/url"

	core "github.com/kkrull/marmot/corerepository"
)

// Registers Git repositories with the meta repo, based upon the address(es) of their remote hosts.
type RegisterRemoteRepositoriesAction struct {
	Source core.RemoteRepositorySource
}

func (cmd *RegisterRemoteRepositoriesAction) Run(remoteUrls []*url.URL) error {
	if addErr := cmd.Source.AddRemotes(remoteUrls); addErr != nil {
		return fmt.Errorf("failed to register remote repositories; %w", addErr)
	} else {
		return nil
	}
}
