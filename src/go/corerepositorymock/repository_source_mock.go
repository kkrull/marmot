package corerepositorymock

import (
	"errors"
	"net/url"

	core "github.com/kkrull/marmot/corerepository"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func NewRepositorySource() *RepositorySource {
	return &RepositorySource{
		RegisterRemoteCalls:  make([]*url.URL, 0),
		RegisterRemoteErrors: make(map[string]error),
		RemoteUrls:           make([]*url.URL, 0),
	}
}

// Mock implementation for testing with RepositorySource.
type RepositorySource struct {
	RegisterRemoteCalls  []*url.URL
	RegisterRemoteErrors map[string]error
	RemoteUrls           []*url.URL
}

func (source *RepositorySource) List() (core.Repositories, error) {
	repositories := make([]core.Repository, len(source.RemoteUrls))
	for i, remoteUrl := range source.RemoteUrls {
		repositories[i] = core.RemoteRepository(remoteUrl)
	}

	return &core.RepositoriesArray{Repositories: repositories}, nil
}

func (source *RepositorySource) RegisterRemote(hostUrl *url.URL) error {
	source.RegisterRemoteCalls = append(source.RegisterRemoteCalls, hostUrl)
	return source.RegisterRemoteErrors[hostUrl.String()]
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

func (source *RepositorySource) RegisterRemoteFails(faultyHref string, errorMsg string) {
	ginkgo.GinkgoHelper()
	source.RegisterRemoteErrors[faultyHref] = errors.New(errorMsg)
}

func (source *RepositorySource) RegisterRemoteNotExpected(unexpectedHref string) {
	ginkgo.GinkgoHelper()

	actualHrefs := make([]string, len(source.RegisterRemoteCalls))
	for i, call := range source.RegisterRemoteCalls {
		actualHrefs[i] = call.String()
	}

	Expect(actualHrefs).NotTo(ContainElement(unexpectedHref))
}
