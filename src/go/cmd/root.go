package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Root command that delegates to other commands
var rootCmd = &cobra.Command{
	Use:   "marmot",
	Short: "Meta Repo Management Tool",
	Long:  `marmot manages a Meta Repository that organizes content in other Git repositories.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Define configuration and flags (persistent and local)
func init() {}

// Add child commands to this root command and set flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
