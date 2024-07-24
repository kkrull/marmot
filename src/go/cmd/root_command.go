package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	debugFlag *bool
	rootCmd   *cobra.Command
)

// Configure the root command with the given I/O and version identifier, then return for use.
func NewRootCommand(stdout io.Writer, stderr io.Writer, version string) (*cobra.Command, error) {
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
	if defaultPath, pathErr := defaultMetaRepoPath(); pathErr != nil {
		return nil, pathErr
	} else {
		rootCmd.PersistentFlags().String("meta-repo", defaultPath, "Meta repo to use")
	}

	// Groups
	rootCmd.AddGroup(&cobra.Group{ID: metaRepoGroup, Title: "Meta Repo Commands"})

	// I/O
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	return rootCmd, nil
}

func defaultMetaRepoPath() (string, error) {
	if homeDir, homeErr := os.UserHomeDir(); homeErr != nil {
		return "", fmt.Errorf("failed to locate home directory; %w", homeErr)
	} else {
		return filepath.Join(homeDir, "meta"), nil
	}
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
	fmt.Printf("Flags:\n")
	fmt.Printf("- debug [%v]: %v\n", rootCmd.Flags().Lookup("debug").DefValue, *debugFlag)
}
