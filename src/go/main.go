package main

import (
	"os"

	"github.com/kkrull/marmot/cmd"
)

func main() {
	rootCmd := cmd.RootCommand("0.0.1")
	if executeErr := rootCmd.Execute(); executeErr != nil {
		os.Exit(1)
	}
}
