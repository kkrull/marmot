package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// Construct a factory to use standard I/O.
func NewCliFactory() *CliFactory {
	return &CliFactory{stdout: os.Stdout, stderr: os.Stderr}
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

func (factory *CliFactory) RootCommand() (*cobra.Command, error) {
	return RootCommand(factory.stdout, factory.stderr, factory.version), nil
}

/* Version configuration */

func (factory *CliFactory) ForExecutable() error {
	if versionPath, versionPathErr := versionFilePath(); versionPathErr != nil {
		return fmt.Errorf("failed to locate version file; %w", versionPathErr)
	} else if version, versionErr := readVersion(versionPath); versionErr != nil {
		return fmt.Errorf("failed to read version from %s; %w", versionPath, versionErr)
	} else {
		factory.version = version
		return nil
	}
}

func readVersion(versionFilename string) (string, error) {
	var versionRaw string
	if versionBytes, readErr := os.ReadFile(versionFilename); readErr != nil {
		return "", fmt.Errorf("failed to read version file %s; %w", versionFilename, readErr)
	} else {
		versionRaw = string(versionBytes)
	}

	if version := strings.TrimSpace(versionRaw); version == "" {
		return "", fmt.Errorf("version <%s> from %s is empty", versionRaw, versionFilename)
	} else {
		return version, nil
	}
}

func versionFilePath() (string, error) {
	if executablePath, executableErr := os.Executable(); executableErr != nil {
		return "", executableErr
	} else {
		programDir := filepath.Dir(executablePath)
		return filepath.Join(programDir, "version"), nil
	}
}
