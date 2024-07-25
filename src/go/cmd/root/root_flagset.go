package cmdroot

import "github.com/spf13/cobra"

// Flag configuration for the root (e.g. top-level) command that dispatches to all other commands.
func RootFlagSet() CommandFlagSet {
	return &rootFlagSet{}
}

// Flags that can be passed to a CLI command.
type CommandFlagSet interface {
	// Add the implemented flags to the given CLI command.
	AddTo(cmd *cobra.Command) error
}

type rootFlagSet struct{}

func (rootFlagSet) AddTo(rootCmd *cobra.Command) error {
	debugFlag.AddTo(rootCmd.PersistentFlags())
	return metaRepoFlag.AddTo(rootCmd.PersistentFlags())
}
