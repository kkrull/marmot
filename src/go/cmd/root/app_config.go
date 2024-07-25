package cmdroot

import (
	"io"
	"net/url"

	"github.com/kkrull/marmot/use"
)

// Abstract factory that creates various application factories from a CLI configuration.
type AppConfig interface {
	AppFactory() use.AppFactory
	Args() []string
	ArgsAsUrls() ([]*url.URL, error)
	Debug() bool
	MetaRepoPath() string
	PrintDebug(writer io.Writer)
}
