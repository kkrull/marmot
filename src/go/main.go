package main

import (
	"fmt"
	"os"

	"github.com/kkrull/marmot/cmd"
)

func main() {
	if version, versionErr := readVersion("version"); versionErr != nil {
		fmt.Fprintln(os.Stderr, versionErr.Error())
		os.Exit(1)
	} else if mainErr := doMain(version); mainErr != nil {
		fmt.Fprintln(os.Stderr, mainErr.Error())
		os.Exit(1)
	}
}

func doMain(version string) error {
	rootCmd := cmd.RootCommand(version)
	return rootCmd.Execute()
}

func readVersion(versionFilename string) (string, error) {
	if versionBytes, readErr := os.ReadFile(versionFilename); readErr != nil {
		return "", fmt.Errorf("failed to read version file %s; %w", versionFilename, readErr)
	} else {
		return string(versionBytes), nil
	}
}
