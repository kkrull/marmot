package cmd

import (
	"os"

	cmdinit "github.com/kkrull/marmot/cmd/init"
	cmdlocal "github.com/kkrull/marmot/cmd/local"
	cmdremote "github.com/kkrull/marmot/cmd/remote"
	cmdshared "github.com/kkrull/marmot/cmd/shared"
	"github.com/spf13/cobra"
)

// TODO KDK: See for example https://github.com/cli/cli/blob/trunk/pkg/cmd/root/root.go

func NewRootCmd() *cobra.Command {
	return rootCmd
}

var rootCmd = &cobra.Command{
	Args: cobra.NoArgs,
	Long: "marmot manages a Meta Repository that organizes content in other (Git) repositories.",
	Use:  "marmot",
	// Uncomment the following line if your bare application has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Short:   "Meta Repo Management Tool",
	Version: "0.0.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cmdinit.NewInitCmd())
	rootCmd.AddCommand(cmdlocal.NewLocalCmd())
	rootCmd.AddCommand(cmdremote.NewRemoteCmd())

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.marmot.yaml)")

	// Cobra also supports local flags, which will only run when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	for _, group := range cmdshared.CommandGroups {
		rootCmd.AddGroup(group.ToCobraGroup())
	}
}
