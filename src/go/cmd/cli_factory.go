package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// Construct a factory to create CLI commands.
func NewCliFactory() *CliFactory {
	return &CliFactory{}
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

/* Version configuration */

func (factory *CliFactory) ForExecutable() (*CliFactory, error) {
	if versionPath, pathErr := versionFilePath(); pathErr != nil {
		return nil, fmt.Errorf("failed to locate version file; %w", pathErr)
	} else if rawVersion, readErr := readVersion(versionPath); readErr != nil {
		return nil, fmt.Errorf("failed to read version from %s; %w", versionPath, readErr)
	} else if version, parseErr := parseVersion(rawVersion); parseErr != nil {
		return nil, fmt.Errorf("failed to parse version from %s; %w", versionPath, parseErr)
	} else {
		factory.version = version
		return factory, nil
	}
}

func parseVersion(versionRaw string) (string, error) {
	if version := strings.TrimSpace(versionRaw); version == "" {
		return "", fmt.Errorf("<%s> is effectively empty", versionRaw)
	} else {
		return version, nil
	}
}

func readVersion(versionFilename string) (string, error) {
	if versionBytes, readErr := os.ReadFile(versionFilename); readErr != nil {
		return "", readErr
	} else {
		return string(versionBytes), nil
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

/* Factory methods */

func (factory *CliFactory) ToRootCobraCommand() (*cobra.Command, error) {
	if rootCmd, err := NewRootCommand(factory.stdout, factory.stderr, factory.version); err != nil {
		return nil, err
	} else {
		NewInitCommand().RegisterWithCobra(rootCmd)
		NewRemoteCommand().RegisterWithCobra(rootCmd)
		return rootCmd, nil
	}
}
