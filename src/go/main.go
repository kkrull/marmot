package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kkrull/marmot/cmd"
)

func main() {
	stdout := os.Stdout
	stderr := os.Stderr
	if version, versionErr := readVersion("version"); versionErr != nil {
		fmt.Fprintln(stderr, versionErr.Error())
		os.Exit(1)
	} else if mainErr := doMain(stdout, stderr, version); mainErr != nil {
		fmt.Fprintln(stderr, mainErr.Error())
		os.Exit(1)
	}
}

func doMain(stdout io.Writer, stderr io.Writer, version string) error {
	rootCmd := cmd.RootCommand(stdout, stderr, version)
	return rootCmd.Execute()
}

func readVersion(versionFilename string) (string, error) {
	if versionBytes, readErr := os.ReadFile(versionFilename); readErr != nil {
		return "", fmt.Errorf("failed to read version file %s; %w", versionFilename, readErr)
	} else {
		return string(versionBytes), nil
	}
}
