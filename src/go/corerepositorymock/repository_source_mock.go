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
		AddLocalCalls:   make([]string, 0),
		AddLocalErrors:  make(map[string]error),
		AddRemoteCalls:  make([]*url.URL, 0),
		AddRemoteErrors: make(map[string]error),
		ListRemoteUrls:  make([]*url.URL, 0),
	}
}

// Mock implementation for testing with RepositorySource.
type RepositorySource struct {
	AddLocalCalls   []string
	AddLocalErrors  map[string]error
	AddRemoteCalls  []*url.URL
	AddRemoteErrors map[string]error
	ListRemoteUrls  []*url.URL
}

/* Local repositories */

func (source *RepositorySource) AddLocal(localPath string) error {
	source.AddLocalCalls = append(source.AddLocalCalls, localPath)
	return source.AddLocalErrors[localPath]
}

func (source *RepositorySource) AddLocalExpected(expectedPaths ...string) {
	ginkgo.GinkgoHelper()
	Expect(source.AddLocalCalls).To(ConsistOf(expectedPaths))
}

func (source *RepositorySource) AddLocalFails(path string, err error) {
	source.AddLocalErrors[path] = err
}

/* Remote repositories */

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
