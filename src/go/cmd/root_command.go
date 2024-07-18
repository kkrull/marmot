package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var (
	debugFlag *bool
	rootCmd   = &cobra.Command{
		Long: `marmot manages a Meta Repository that organizes content in other (Git) repositories.`,
		Run: func(cmd *cobra.Command, args []string) {
			if *debugFlag {
				printDebug()
			}
		},
		Short: "Meta Repo Management Tool",
		Use:   "marmot",
	}
)

// Configure the root command with the given I/O and version identifier, then return for use.
func RootCommand(stdout io.Writer, stderr io.Writer, version string) *cobra.Command {
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	rootCmd.Version = version
	return rootCmd
}

func init() {
	debugFlag = rootCmd.PersistentFlags().Bool("debug", false, "print CLI debugging information")
	rootCmd.PersistentFlags().Lookup("debug").Hidden = true
}

/* Pseudo-commands */

func printDebug() {
	fmt.Printf("--debug: %v\n", *debugFlag)
}
