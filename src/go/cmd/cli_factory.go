package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

// Construct a factory to create CLI commands.
func NewCliFactory(version string) *CliFactory {
	return &CliFactory{version: version}
}

// Creates commands for the Command Line Interface (CLI).
type CliFactory struct {
	stdout  io.Writer
	stderr  io.Writer
	version string
}

func (factory *CliFactory) WithStdIO(stdout io.Writer, stderr io.Writer) *CliFactory {
	factory.stdout = stdout
	factory.stderr = stderr
	return factory
}

/* Factory methods */

func (factory *CliFactory) ToRootCobraCommand() *cobra.Command {
	return NewRootCommand(factory.stdout, factory.stderr, factory.version).ToCobraCommand()
}
