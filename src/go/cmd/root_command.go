package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

// Configure the root command with the given I/O and version identifier, then return for use.
func NewRootCommand(stdout io.Writer, stderr io.Writer, version string) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Long:    "marmot manages a Meta Repository that organizes content in other (Git) repositories.",
		RunE:    runRoot,
		Short:   "Meta Repo Management Tool",
		Use:     "marmot",
		Version: version,
	}

	AddFlags(rootCmd)
	addGroups(rootCmd)
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	return rootCmd, nil
}

func runRoot(cobraCmd *cobra.Command, args []string) error {
	if config, parseErr := ParseFlags(cobraCmd); parseErr != nil {
		return parseErr
	} else if config.Debug {
		config.PrintDebug(cobraCmd.OutOrStdout())
		return nil
	} else if len(args) == 0 {
		return cobraCmd.Help()
	} else {
		return nil
	}
}

/* Child commands */

const (
	metaRepoGroup = "meta-repo"
)

func addGroups(cobraCmd *cobra.Command) {
	cobraCmd.AddGroup(&cobra.Group{ID: metaRepoGroup, Title: "Meta Repo Commands"})
}

func AddMetaRepoCommand(parent *cobra.Command, child cobra.Command) {
	child.GroupID = metaRepoGroup
	parent.AddCommand(&child)
}
