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
		AddRemoteCalls:  make([]*url.URL, 0),
		AddRemoteErrors: make(map[string]error),
		ListRemoteUrls:  make([]*url.URL, 0),
	}
}

// Mock implementation for testing with RepositorySource.
type RepositorySource struct {
	AddRemoteCalls  []*url.URL
	AddRemoteErrors map[string]error
	ListRemoteUrls  []*url.URL
}

func (source *RepositorySource) AddRemote(hostUrl *url.URL) error {
	source.AddRemoteCalls = append(source.AddRemoteCalls, hostUrl)
	return source.AddRemoteErrors[hostUrl.String()]
}

func (source *RepositorySource) AddRemoteExpected(expectedHref string) {
	ginkgo.GinkgoHelper()
	actualHrefs := source.addRemoteHrefs()
	Expect(actualHrefs).To(ContainElement(expectedHref))
}

func (source *RepositorySource) AddRemoteFails(faultyHref string, errorMsg string) {
	ginkgo.GinkgoHelper()
	source.AddRemoteErrors[faultyHref] = errors.New(errorMsg)
}

func (source *RepositorySource) AddRemoteNotExpected(unexpectedHref string) {
	ginkgo.GinkgoHelper()
	actualHrefs := source.addRemoteHrefs()
	Expect(actualHrefs).NotTo(ContainElement(unexpectedHref))
}

func (source *RepositorySource) addRemoteHrefs() []string {
	actualHrefs := make([]string, len(source.AddRemoteCalls))
	for i, call := range source.AddRemoteCalls {
		actualHrefs[i] = call.String()
	}

	return actualHrefs
}

func (source *RepositorySource) ListRemote() (core.Repositories, error) {
	repositories := make([]core.Repository, len(source.ListRemoteUrls))
	for i, remoteUrl := range source.ListRemoteUrls {
		repositories[i] = core.RemoteRepository(remoteUrl)
	}

	return &core.RepositoriesArray{Repositories: repositories}, nil
}
