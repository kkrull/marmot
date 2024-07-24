package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

var (
	rootCmd *cobra.Command
)

// Configure the root command with the given I/O and version identifier, then return for use.
func NewRootCommand(stdout io.Writer, stderr io.Writer, version string) (*cobra.Command, error) {
	rootCmd = &cobra.Command{
		Long: "marmot manages a Meta Repository that organizes content in other (Git) repositories.",
		RunE: func(cmd *cobra.Command, args []string) error {
			config := ParseFlags(cmd)
			if config.Debug() {
				config.PrintDebug(stdout)
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
	AddFlags(rootCmd)

	// Groups
	rootCmd.AddGroup(&cobra.Group{ID: metaRepoGroup, Title: "Meta Repo Commands"})

	// I/O
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	return rootCmd, nil
}

/* Child commands */

const (
	metaRepoGroup = "meta-repo"
)

func AddMetaRepoCommand(child cobra.Command) {
	child.GroupID = metaRepoGroup
	rootCmd.AddCommand(&child)
}
