package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	debugFlag *bool
)

// Root command that delegates to other commands
func RootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Long: `marmot manages a Meta Repository that organizes content in other (Git) repositories.`,
		Run: func(cmd *cobra.Command, args []string) {
			if *debugFlag {
				printDebug()
			}
		},
		Short:   "Meta Repo Management Tool",
		Use:     "marmot",
		Version: "0.0.1",
	}

	// Add flags (persistent and local)
	debugFlag = rootCmd.PersistentFlags().Bool("debug", false, "print CLI debugging information")
	rootCmd.PersistentFlags().Lookup("debug").Hidden = true

	// Add child (sub-)commands
	return rootCmd
}

/* Pseudo-commands */

func printDebug() {
	fmt.Printf("--debug: %v\n", *debugFlag)
}
