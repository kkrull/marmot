package userepository

import (
	"net/url"

	core "github.com/kkrull/marmot/corerepository"
)

// Registers Git repositories with the meta repo, based upon the address(es) of their remote hosts.
type RegisterRemoteRepositoriesCommand struct {
	Source core.RepositorySource
}

func (cmd *RegisterRemoteRepositoriesCommand) Run(remoteUrls []url.URL) error {
	_ = cmd.Source.RegisterRemote(remoteUrls[0])
	return nil
}
