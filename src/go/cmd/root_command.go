package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionFlag *bool
)

// Root command that delegates to other commands
func RootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Long: `marmot manages a Meta Repository that organizes content in other (Git) repositories.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("--version: %v\n", *versionFlag)
		},
		Short: "Meta Repo Management Tool",
		Use:   "marmot",
	}

	// Add flags (persistent and local)
	versionFlag = rootCmd.LocalFlags().Bool("version", false, "print the marmot suite version")

	// Add child (sub-)commands
	return rootCmd
}
