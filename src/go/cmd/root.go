package cmd

import (
	"github.com/spf13/cobra"
)

// Root command that delegates to other commands
var rootCmd = &cobra.Command{
	Long:  `marmot manages a Meta Repository that organizes content in other (Git) repositories.`,
	Short: "Meta Repo Management Tool",
	Use:   "marmot",
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Define configuration and flags (persistent and local)
func init() {}

// Add child commands to this root command and set flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}
