package mainfactory

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/kkrull/marmot/cmd"
	"github.com/kkrull/marmot/cmdinit"
	"github.com/spf13/cobra"
)

// Construct a factory to create CLI commands.
func NewCliFactory(appFactory *AppFactory) *CliFactory {
	return &CliFactory{appFactory: appFactory}
}

// Creates commands for the Command Line Interface (CLI).
type CliFactory struct {
	appFactory *AppFactory
	stdout     io.Writer
	stderr     io.Writer
	version    string
}

func (cliFactory *CliFactory) WithStdIO(stdout io.Writer, stderr io.Writer) *CliFactory {
	cliFactory.stdout = stdout
	cliFactory.stderr = stderr
	return cliFactory
}

/* Factory methods */

func (cliFactory *CliFactory) CommandTree() (*cobra.Command, error) {
	rootCmd := cmd.NewRootCommand(cliFactory.stdout, cliFactory.stderr, cliFactory.version)
	if initAppCmd, appFactoryErr := cliFactory.appFactory.InitCommand(); appFactoryErr != nil {
		return nil, appFactoryErr
	} else if metaRepoPath, pathErr := defaultMetaRepoPath(); pathErr != nil {
		return nil, pathErr
	} else {
		initCliCmd := cmdinit.NewInitCommand(initAppCmd, metaRepoPath)
		cmd.AddMetaRepoCommand(initCliCmd.ToCobraCommand())
		return rootCmd, nil
	}
}

/* Version configuration */

func (cliFactory *CliFactory) ForExecutable() (*CliFactory, error) {
	if versionPath, pathErr := versionFilePath(); pathErr != nil {
		return nil, fmt.Errorf("failed to locate version file; %w", pathErr)
	} else if rawVersion, readErr := readVersion(versionPath); readErr != nil {
		return nil, fmt.Errorf("failed to read version from %s; %w", versionPath, readErr)
	} else if version, parseErr := parseVersion(rawVersion); parseErr != nil {
		return nil, fmt.Errorf("failed to parse version from %s; %w", versionPath, parseErr)
	} else {
		cliFactory.version = version
		return cliFactory, nil
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