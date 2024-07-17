package userepository

import (
	"fmt"
	"net/url"

	core "github.com/kkrull/marmot/corerepository"
)

// Registers Git repositories with the meta repo, based upon the address(es) of their remote hosts.
type RegisterRemoteRepositoriesCommand struct {
	Source core.RepositorySource
}

func (cmd *RegisterRemoteRepositoriesCommand) Run(remoteUrls []*url.URL) error {
	for _, remoteUrl := range remoteUrls {
		if registerErr := cmd.Source.RegisterRemote(remoteUrl); registerErr != nil {
			return fmt.Errorf("failed to register %s; %w", remoteUrl.String(), registerErr)
		}
	}

	return nil
}
