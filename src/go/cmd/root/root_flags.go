package cmdroot

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Flag configuration for the root (e.g. top-level) command that dispatches to all other commands.
func RootFlagSet() CommandFlags {
	return &rootFlags{}
}

// Flags that can be passed to a CLI command.
type CommandFlags interface {
	// Add the implemented flags to the given CLI command.
	AddTo(cmd *cobra.Command) error
}

type rootFlags struct{}

func (rootFlags) AddTo(rootCmd *cobra.Command) error {
	addDebugFlag(rootCmd.PersistentFlags())
	if metaRepoErr := addMetaRepoFlag(rootCmd.PersistentFlags()); metaRepoErr != nil {
		return metaRepoErr
	}

	return nil
}

func addDebugFlag(flags *pflag.FlagSet) {
	flags.Bool("debug", false, "print CLI debugging information")
	flags.Lookup("debug").Hidden = true
}

func addMetaRepoFlag(flags *pflag.FlagSet) error {
	if homeDir, homeErr := os.UserHomeDir(); homeErr != nil {
		return fmt.Errorf("failed to locate home directory; %w", homeErr)
	} else {
		flags.String("meta-repo", filepath.Join(homeDir, "meta"), "Meta repo to use")
		return nil
	}
}
