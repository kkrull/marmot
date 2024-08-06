package cmd

import (
	"io"

	cmdroot "github.com/kkrull/marmot/cmd/root"
	"github.com/spf13/cobra"
)

// Construct a root CLI command with the given I/O and version identifier.
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

// Map to a command that runs on Cobra.
func (root rootCliCommand) ToCobraCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Args:    cobra.NoArgs,
		Long:    "marmot manages a Meta Repository that organizes content in other (Git) repositories.",
		RunE:    runRootCobra,
		Short:   "Meta Repo Management Tool",
		Use:     "marmot",
		Version: root.version,
	}

	cmdroot.RootFlagSet().AddTo(rootCmd)
	for _, group := range commandGroups {
		rootCmd.AddGroup(group.toCobraGroup())
	}

	rootCmd.AddCommand(
		NewInitCommand().ToCobraCommand(),
		NewRemoteCommand().ToCobraCommand(),
	)

	rootCmd.SetOut(root.stdout)
	rootCmd.SetErr(root.stderr)

	return rootCmd
}

func runRootCobra(cli *cobra.Command, args []string) error {
	if parser, newErr := cmdroot.RootConfigParser(); newErr != nil {
		return newErr
	} else if config, parseErr := parser.Parse(cli.Flags(), args); parseErr != nil {
		return parseErr
	} else if config.Debug() {
		config.PrintDebug(cli.OutOrStdout())
		return nil
	} else if len(args) == 0 {
		return cli.Help()
	} else {
		// Run the command named in the arguments
		return nil
	}
}
