package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	debugFlag   *bool
	versionFlag *bool
)

// Root command that delegates to other commands
func RootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Long: `marmot manages a Meta Repository that organizes content in other (Git) repositories.`,
		Run: func(cmd *cobra.Command, args []string) {
			if *debugFlag {
				printDebug()
			} else if *versionFlag {
				printVersion()
			}
		},
		Short: "Meta Repo Management Tool",
		Use:   "marmot",
	}

	// Add flags (persistent and local)
	debugFlag = rootCmd.PersistentFlags().Bool("debug", false, "print CLI debugging information")
	rootCmd.PersistentFlags().Lookup("debug").Hidden = true

	versionFlag = rootCmd.PersistentFlags().Bool("version", false, "print the marmot suite version")

	// Add child (sub-)commands
	return rootCmd
}

/* Pseudo-commands */

func printDebug() {
	fmt.Printf("--debug: %v\n", *debugFlag)
	fmt.Printf("--version: %v\n", *versionFlag)
}

func printVersion() {
	fmt.Printf("marmot (go) 0.0.1")
}
