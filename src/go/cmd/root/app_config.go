package cmdroot

import (
	"io"

	"github.com/kkrull/marmot/use"
)

// Abstract factory that creates various application factories from a CLI configuration.
type AppConfig interface {
	AppFactory() use.AppFactory
	Debug() bool
	MetaRepoPath() string
	PrintDebug(writer io.Writer)
}
