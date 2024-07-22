package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var (
	debugFlag *bool
	rootCmd   = &cobra.Command{
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
		Short: "Meta Repo Management Tool",
		Use:   "marmot [--help|--version]",
	}
)

// Configure the root command with the given I/O and version identifier, then return for use.
func NewRootCommand(stdout io.Writer, stderr io.Writer, version string) *cobra.Command {
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	rootCmd.Version = version
	return rootCmd
}

/* Child commands */

const (
	metaRepoGroup = "meta-repo"
)

func AddCommandToRoot(child *cobra.Command) { //TODO KDK: Add group to encapsulate constant
	rootCmd.AddCommand(child)
}

/* Configuration */

func init() {
	initFlags()
	initGroups()
}

func initFlags() {
	debugFlag = rootCmd.PersistentFlags().Bool("debug", false, "print CLI debugging information")
	rootCmd.PersistentFlags().Lookup("debug").Hidden = true
}

func initGroups() {
	rootCmd.AddGroup(&cobra.Group{ID: "meta-repo", Title: "Meta Repo Commands"})
}

/* Pseudo-commands */

func printDebug() {
	fmt.Printf("--debug: %v\n", *debugFlag)
}
