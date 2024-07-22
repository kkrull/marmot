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

func (factory *CliFactory) NewRootCommand() (*cobra.Command, error) {
	return RootCommand(factory.stdout, factory.stderr, factory.version), nil
}

/* Version configuration */

func (factory *CliFactory) ForExecutable() error {
	if versionPath, pathErr := versionFilePath(); pathErr != nil {
		return fmt.Errorf("failed to locate version file; %w", pathErr)
	} else if rawVersion, readErr := readVersion(versionPath); readErr != nil {
		return fmt.Errorf("failed to read version from %s; %w", versionPath, readErr)
	} else if version, parseErr := parseVersion(rawVersion); parseErr != nil {
		return fmt.Errorf("failed to parse version from %s; %w", versionPath, parseErr)
	} else {
		factory.version = version
		return nil
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
