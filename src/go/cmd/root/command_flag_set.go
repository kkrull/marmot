package cmdroot

import (
	"fmt"
	"io"

	"github.com/kkrull/marmot/use"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Flag configuration for a single CLI command
type CommandFlagSet interface {
	// Add the implemented flags to the given CLI command
	AddTo(cmd *cobra.Command) error

	// Parse application configuration from flags passed to the CLI
	ParseAppConfig(flags *pflag.FlagSet, args []string) (AppConfig, error)
}

// Application configuration derived from flags passed to the CLI.
type FlagAppConfig struct {
	appFactory   use.AppFactory
	debug        bool
	flagSet      *pflag.FlagSet
	metaRepoPath string
}

func (config FlagAppConfig) AppFactory() use.AppFactory { return config.appFactory }
func (config FlagAppConfig) Debug() bool                { return config.debug }
func (config FlagAppConfig) MetaRepoPath() string       { return config.metaRepoPath }
func (config FlagAppConfig) PrintDebug(writer io.Writer) {
	fmt.Fprintf(writer, "Flags:\n")

	debugFlag := config.flagSet.Lookup("debug")
	fmt.Fprintf(writer, "- debug [%v]: %v\n", debugFlag.DefValue, debugFlag.Value)

	metaRepoFlag := config.flagSet.Lookup("meta-repo")
	fmt.Fprintf(writer, "- meta-repo [%v]: %v\n", metaRepoFlag.DefValue, metaRepoFlag.Value)
}
