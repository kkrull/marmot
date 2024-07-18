package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var (
	debugFlag *bool
)

// Root command that delegates to other commands
func RootCommand(stdout io.Writer, stderr io.Writer, version string) *cobra.Command {
	var rootCmd = &cobra.Command{
		Long: `marmot manages a Meta Repository that organizes content in other (Git) repositories.`,
		Run: func(cmd *cobra.Command, args []string) {
			if *debugFlag {
				printDebug()
			}
		},
		Short:   "Meta Repo Management Tool",
		Use:     "marmot",
		Version: version,
	}

	// Add flags
	debugFlag = rootCmd.PersistentFlags().Bool("debug", false, "print CLI debugging information")
	rootCmd.PersistentFlags().Lookup("debug").Hidden = true

	// Add child (sub-)commands

	// Configure I/O
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	return rootCmd
}

/* Pseudo-commands */

func printDebug() {
	fmt.Printf("--debug: %v\n", *debugFlag)
}
