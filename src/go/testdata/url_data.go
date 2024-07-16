package testdata

import (
	"net/url"

	"github.com/kkrull/marmot/expect"
	"github.com/onsi/ginkgo/v2"
)

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
