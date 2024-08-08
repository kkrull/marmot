package corerepositorymock

import (
	"errors"
	"net/url"

	core "github.com/kkrull/marmot/corerepository"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Construct a test double of RepositorySource.
func NewRepositorySource() *RepositorySource {
	return &RepositorySource{
		addLocalCalls:  make([]string, 0),
		addLocalErrors: make(map[string]error),
		addRemoteCalls: make([]*url.URL, 0),
		ListLocalPaths: make([]string, 0),
		ListRemoteUrls: make([]*url.URL, 0),
	}
}

// Mock implementation for testing with RepositorySource.
type RepositorySource struct {
	addLocalCalls   []string
	addLocalErrors  map[string]error
	addLocalsError  error
	addRemoteCalls  []*url.URL
	addRemotesError error
	ListLocalPaths  []string
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

func (source *RepositorySource) AddLocals(localPaths []string) error {
	source.addLocalCalls = append(source.addLocalCalls, localPaths...)
	return source.addLocalsError
}

func (source *RepositorySource) ListLocal() (core.Repositories, error) {
	repositories := make([]core.Repository, len(source.ListLocalPaths))
	for i, localPath := range source.ListLocalPaths {
		repositories[i] = core.LocalRepository(localPath)
	}

	return core.SomeRepositories(repositories), nil
}

/* Remote repositories */

func (source *RepositorySource) AddRemotes(hostUrls []*url.URL) error {
	source.addRemoteCalls = append(source.addRemoteCalls, hostUrls...)
	return source.addRemotesError
}

func (source *RepositorySource) AddRemotesExpected(expectedHrefs ...string) {
	ginkgo.GinkgoHelper()
	actualHrefs := source.addRemoteHrefs()
	Expect(actualHrefs).To(ConsistOf(expectedHrefs))
}

func (source *RepositorySource) AddRemotesFails(errorMsg string) {
	source.addRemotesError = errors.New(errorMsg)
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

	return core.SomeRepositories(repositories), nil
}
