package cmd

import (
	cmdinit "github.com/kkrull/marmot/cmd/init"
	cmdlocal "github.com/kkrull/marmot/cmd/local"
	cmdremote "github.com/kkrull/marmot/cmd/remote"
	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

func NewRootCmd(metaRepoDefault string, version string) *cobra.Command {
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

	cmdshared.FlagSet().AddTo(rootCmd)
	cmdshared.AddGroups(rootCmd)
	return rootCmd
}
