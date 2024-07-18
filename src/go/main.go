package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/kkrull/marmot/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var (
		stdout  io.Writer = os.Stdout
		stderr  io.Writer = os.Stderr
		rootCmd *cobra.Command
	)

	if versionFilename, pathErr := filepath.Abs("version"); pathErr != nil {
		fmt.Fprintf(stderr, "failed to locate version file; %s\n", pathErr.Error())
		os.Exit(1)
	} else if version, versionErr := readVersion(versionFilename); versionErr != nil {
		fmt.Fprintf(stderr, "failed to read version; %s\n", versionErr.Error())
		os.Exit(1)
	} else {
		rootCmd = cmd.RootCommand(stdout, stderr, version)
	}

	if executeErr := rootCmd.Execute(); executeErr != nil {
		fmt.Fprintln(stderr, executeErr.Error())
		os.Exit(1)
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
