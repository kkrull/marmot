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
		addLocalCalls:   make([]string, 0),
		addLocalErrors:  make(map[string]error),
		addRemoteCalls:  make([]*url.URL, 0),
		addRemoteErrors: make(map[string]error),
		ListRemoteUrls:  make([]*url.URL, 0),
	}
}

// Mock implementation for testing with RepositorySource.
type RepositorySource struct {
	addLocalCalls   []string
	addLocalErrors  map[string]error
	addRemoteCalls  []*url.URL
	addRemoteErrors map[string]error
	ListRemoteUrls  []*url.URL
}

/* Local repositories */

func (source *RepositorySource) AddLocal(localPath string) error {
	source.addLocalCalls = append(source.addLocalCalls, localPath)
	return source.addLocalErrors[localPath]
}

func (source *RepositorySource) AddLocalExpected(expectedPaths ...string) {
	ginkgo.GinkgoHelper()
	Expect(source.addLocalCalls).To(ConsistOf(expectedPaths))
}

func (source *RepositorySource) AddLocalFails(path string, err error) {
	source.addLocalErrors[path] = err
}

/* Remote repositories */

func (source *RepositorySource) AddRemote(hostUrl *url.URL) error {
	source.addRemoteCalls = append(source.addRemoteCalls, hostUrl)
	return source.addRemoteErrors[hostUrl.String()]
}

func (source *RepositorySource) AddRemoteExpected(expectedHref string) {
	ginkgo.GinkgoHelper()
	actualHrefs := source.addRemoteHrefs()
	Expect(actualHrefs).To(ContainElement(expectedHref))
}

func (source *RepositorySource) AddRemoteFails(faultyHref string, errorMsg string) {
	source.addRemoteErrors[faultyHref] = errors.New(errorMsg)
}

func (source *RepositorySource) AddRemoteNotExpected(unexpectedHref string) {
	ginkgo.GinkgoHelper()
	actualHrefs := source.addRemoteHrefs()
	Expect(actualHrefs).NotTo(ContainElement(unexpectedHref))
}

func (source *RepositorySource) addRemoteHrefs() []string {
	actualHrefs := make([]string, len(source.addRemoteCalls))
	for i, call := range source.addRemoteCalls {
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
