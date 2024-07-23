package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var (
	debugFlag *bool
	rootCmd   *cobra.Command
)

// Configure the root command with the given I/O and version identifier, then return for use.
func NewRootCommand(stdout io.Writer, stderr io.Writer, version string) *cobra.Command {
	rootCmd = &cobra.Command{
		Long: "marmot manages a Meta Repository that organizes content in other (Git) repositories.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if *debugFlag {
				printDebug()
				return nil
			} else if len(args) == 0 {
				return cmd.Help()
			} else {
				return nil
			}
		},
		Short:   "Meta Repo Management Tool",
		Use:     "marmot [--help|--version]",
		Version: version,
	}

	// Flags
	debugFlag = rootCmd.PersistentFlags().Bool("debug", false, "print CLI debugging information")
	rootCmd.PersistentFlags().Lookup("debug").Hidden = true

	// Groups
	rootCmd.AddGroup(&cobra.Group{ID: metaRepoGroup, Title: "Meta Repo Commands"})

	// I/O
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	return rootCmd
}

/* Child commands */

const (
	metaRepoGroup = "meta-repo"
)

func AddMetaRepoCommand(child cobra.Command) {
	child.GroupID = metaRepoGroup
	rootCmd.AddCommand(&child)
}

/* Pseudo-commands */

func printDebug() {
	fmt.Printf("--debug: %v\n", *debugFlag)
}
