package cmd

import (
	cmdinit "github.com/kkrull/marmot/cmd/init"
	cmdlocal "github.com/kkrull/marmot/cmd/local"
	cmdremote "github.com/kkrull/marmot/cmd/remote"
	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

var (
	metaRepoGroup   cmdshared.CommandGroup   = cmdshared.NewCommandGroup("meta-repo-group", "Meta Repo Commands")
	repositoryGroup cmdshared.CommandGroup   = cmdshared.NewCommandGroup("repository-group", "Repository Commands")
	groups          []cmdshared.CommandGroup = []cmdshared.CommandGroup{
		metaRepoGroup,
		repositoryGroup,
	}
)

func NewRootCmd(metaRepoDefault string, version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		Long:    "marmot manages a Meta Repository that organizes content in other (Git) repositories.",
		Use:     "marmot",
		Short:   "Meta Repo Management Tool",
		Version: version,
	}

	rootCmd.AddCommand(cmdinit.NewInitCmd(metaRepoGroup))
	rootCmd.AddCommand(cmdlocal.NewLocalCmd(repositoryGroup))
	rootCmd.AddCommand(cmdremote.NewRemoteCmd(repositoryGroup))

	cmdshared.FlagSet().AddTo(rootCmd)
	for _, group := range groups {
		rootCmd.AddGroup(group.ToCobraGroup())
	}

	return rootCmd
}
