package cmd

import (
	"github.com/spf13/cobra"
)

// Root command that delegates to other commands
func RootCommand() *cobra.Command {
	//Omit .Run, which is for sub-commands
	var rootCmd = &cobra.Command{
		Long:  `marmot manages a Meta Repository that organizes content in other (Git) repositories.`,
		Short: "Meta Repo Management Tool",
		Use:   "marmot",
	}

	// Add flags (persistent and local) and child (sub-)commands
	return rootCmd
}
