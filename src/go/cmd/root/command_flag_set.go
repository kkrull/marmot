package cmdroot

import (
	"fmt"
	"io"
	"net/url"

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
	args         []string
	debug        bool
	flagSet      *pflag.FlagSet
	metaRepoPath string
}

func (config FlagAppConfig) AppFactory() use.AppFactory { return config.appFactory }
func (config FlagAppConfig) Args() []string             { return config.args }

func (config FlagAppConfig) ArgsAsUrls() ([]*url.URL, error) {
	urls := make([]*url.URL, len(config.args))
	for i, rawArg := range config.args {
		if urlArg, parseErr := url.Parse(rawArg); parseErr != nil {
			return nil, fmt.Errorf("url expected: %s; %w", rawArg, parseErr)
		} else {
			urls[i] = urlArg
		}
	}

	return urls, nil
}

func (config FlagAppConfig) Debug() bool          { return config.debug }
func (config FlagAppConfig) MetaRepoPath() string { return config.metaRepoPath }
func (config FlagAppConfig) PrintDebug(writer io.Writer) {
	for i, arg := range config.args {
		fmt.Fprintf(writer, "[%d]: %s\n", i, arg)
	}

	debugFlag := config.flagSet.Lookup("debug")
	fmt.Fprintf(writer, "--debug [%v]: %v\n", debugFlag.DefValue, debugFlag.Value)

	metaRepoFlag := config.flagSet.Lookup("meta-repo")
	fmt.Fprintf(writer, "--meta-repo [%v]: %v\n", metaRepoFlag.DefValue, metaRepoFlag.Value)
}
