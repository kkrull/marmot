package testsupportdata

import (
	"net/url"

	expect "github.com/kkrull/marmot/testsupportexpect"
	"github.com/onsi/ginkgo/v2"
)

// Parse URL from valid href string, or fail
func NewURL(rawUrl string) *url.URL {
	ginkgo.GinkgoHelper()
	return expect.NoError(url.Parse(rawUrl))
}

// Parse URLs from valid href strings, or fail
func NewURLs(rawUrls ...string) []*url.URL {
	ginkgo.GinkgoHelper()
	parsedUrls := make([]*url.URL, len(rawUrls))
	for i, rawUrl := range rawUrls {
		parsedUrl := expect.NoError(url.Parse(rawUrl))
		parsedUrls[i] = parsedUrl
	}

	return parsedUrls
}
