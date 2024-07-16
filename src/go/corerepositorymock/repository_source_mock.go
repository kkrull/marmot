package corerepositorymock

import (
	"net/url"

	core "github.com/kkrull/marmot/corerepository"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Mock implementation for testing with RepositorySource.
type RepositorySource struct {
	Names               []string
	RegisterRemoteCalls []*url.URL
}

func (source *RepositorySource) List() (core.Repositories, error) {
	repositories := make([]core.Repository, len(source.Names))
	for i, name := range source.Names {
		repositories[i] = core.Repository{Name: name}
	}

	return &core.RepositoriesArray{Repositories: repositories}, nil
}

func (source *RepositorySource) RegisterRemote(hostUrl *url.URL) error {
	source.RegisterRemoteCalls = append(source.RegisterRemoteCalls, hostUrl)
	return nil
}

func (source *RepositorySource) RegisterRemoteExpected(expectedHref string) {
	ginkgo.GinkgoHelper()

	actualHrefs := make([]string, len(source.RegisterRemoteCalls))
	for i, call := range source.RegisterRemoteCalls {
		actualHrefs[i] = call.String()
		if call.String() == expectedHref {
			return
		}
	}

	Expect(actualHrefs).To(ContainElement(expectedHref))
}