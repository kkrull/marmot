package userepository

import (
	"net/url"
)

// Registers Git repositories with the meta repo, based upon the address(es) of their remote hosts.
type RegisterRemoteRepositoriesCommand struct{}

func (cmd *RegisterRemoteRepositoriesCommand) Run(remoteUrls []url.URL) error {
	return nil
}
