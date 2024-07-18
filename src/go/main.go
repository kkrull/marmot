package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kkrull/marmot/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var (
		stdout io.Writer = os.Stdout
		stderr io.Writer = os.Stderr
	)

	if rootCmd, initErr := newRootCommand(stdout, stderr); initErr != nil {
		fmt.Fprintln(stderr, initErr.Error())
		os.Exit(1)
	} else if executeErr := rootCmd.Execute(); executeErr != nil {
		fmt.Fprintln(stderr, executeErr.Error())
		os.Exit(1)
	}
}

func newRootCommand(stdout io.Writer, stderr io.Writer) (*cobra.Command, error) {
	if version, versionErr := readVersion("version"); versionErr != nil {
		return nil, fmt.Errorf("failed to read version; %w", versionErr)
	} else {
		return cmd.RootCommand(stdout, stderr, version), nil
	}
}

func readVersion(versionFilename string) (string, error) {
	var version string
	if versionBytes, readErr := os.ReadFile(versionFilename); readErr != nil {
		return "", fmt.Errorf("failed to read version file %s; %w", versionFilename, readErr)
	} else {
		version = string(versionBytes)
	}

	trimmed := strings.TrimSpace(version)
	if trimmed == "" {
		return "", fmt.Errorf("version <%s> from %s is empty", version, versionFilename)
	} else {
		return trimmed, nil
	}
}
