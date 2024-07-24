package cmd

import (
	"io"

	"github.com/kkrull/marmot/use"
)

// Application configuration that doesn't have anything particular to do with the CLI.
type AppConfig interface {
	AppFactory() *use.AppFactory
	Debug() bool
	MetaRepoPath() string
	PrintDebug(writer io.Writer)
}
