package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

// Configure the root command with the given I/O and version identifier, then return for use.
func NewRootCommand(stdout io.Writer, stderr io.Writer, version string) *rootCliCommand {
	return &rootCliCommand{
		stderr:  stderr,
		stdout:  stdout,
		version: version,
	}
}

type rootCliCommand struct {
	stderr  io.Writer
	stdout  io.Writer
	version string
}

func (root rootCliCommand) ToCobraCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Long:    "marmot manages a Meta Repository that organizes content in other (Git) repositories.",
		RunE:    runRoot,
		Short:   "Meta Repo Management Tool",
		Use:     "marmot",
		Version: root.version,
	}

	RootFlagSet().AddTo(rootCmd)
	for _, group := range commandGroups {
		rootCmd.AddGroup(group.toCobraGroup())
	}

	rootCmd.AddCommand(
		NewInitCommand().toCobraCommand(),
		NewRemoteCommand().toCobraCommand(),
	)

	rootCmd.SetOut(root.stdout)
	rootCmd.SetErr(root.stderr)

	return rootCmd
}

func runRoot(cobraCmd *cobra.Command, args []string) error {
	if config, parseErr := RootFlagSet().ParseAppConfig(cobraCmd.Flags(), args); parseErr != nil {
		return parseErr
	} else if config.Debug() {
		config.PrintDebug(cobraCmd.OutOrStdout())
		return nil
	} else if len(args) == 0 {
		return cobraCmd.Help()
	} else {
		return nil
	}
}
