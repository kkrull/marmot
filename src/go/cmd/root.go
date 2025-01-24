package cmd

import (
	cmdinit "github.com/kkrull/marmot/cmd/init"
	cmdlocal "github.com/kkrull/marmot/cmd/local"
	cmdremote "github.com/kkrull/marmot/cmd/remote"
	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

// TODO KDK: See for example https://github.com/cli/cli/blob/trunk/pkg/cmd/root/root.go
// https://github.com/cli/cli/blob/trunk/pkg/cmd/pr/pr.go

func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		Long:    "marmot manages a Meta Repository that organizes content in other (Git) repositories.",
		Use:     "marmot",
		Short:   "Meta Repo Management Tool",
		Version: version,
	}

	rootCmd.AddCommand(cmdinit.NewInitCmd())
	rootCmd.AddCommand(cmdlocal.NewLocalCmd())
	rootCmd.AddCommand(cmdremote.NewRemoteCmd())

	// Cobra supports persistent flags, which, if defined here, will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.marmot.yaml)")

	for _, group := range cmdshared.CommandGroups {
		rootCmd.AddGroup(group.ToCobraGroup())
	}

	return rootCmd
}
