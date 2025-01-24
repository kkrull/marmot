package cmd

import (
	cmdinit "github.com/kkrull/marmot/cmd/init"
	cmdlocal "github.com/kkrull/marmot/cmd/local"
	cmdremote "github.com/kkrull/marmot/cmd/remote"
	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

var (
	metaRepoGroup   cmdshared.CommandGroup = cmdshared.NewCommandGroup("meta-repo-group", "Meta Repo Commands")
	repositoryGroup cmdshared.CommandGroup = cmdshared.NewCommandGroup("repository-group", "Repository Commands")
)

func NewRootCmd(metaRepoDefault string, version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		Long:    "marmot manages a Meta Repository that organizes content in other (Git) repositories.",
		Use:     "marmot",
		Short:   "Meta Repo Management Tool",
		Version: version,
	}

	cmdshared.FlagSet().AddTo(rootCmd)
	rootCmd.AddGroup(
		metaRepoGroup.ToCobraGroup(),
		repositoryGroup.ToCobraGroup(),
	)

	rootCmd.AddCommand(
		cmdinit.NewInitCmd(metaRepoGroup),
		cmdlocal.NewLocalCmd(repositoryGroup),
		cmdremote.NewRemoteCmd(repositoryGroup),
	)

	return rootCmd
}
