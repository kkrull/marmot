package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kkrull/marmot/use"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

/* Root command flags */

// Flag configuration for the root (e.g. top-level) command that dispatches to all other commands.
func RootFlagSet() CommandFlagSet {
	return &rootFlagSet{}
}

type rootFlagSet struct{}

func (rootFlagSet) AddTo(rootCmd *cobra.Command) error {
	addDebugFlag(rootCmd.PersistentFlags())
	if metaRepoErr := addMetaRepoFlag(rootCmd.PersistentFlags()); metaRepoErr != nil {
		return metaRepoErr
	}

	return nil
}

func addDebugFlag(flags *pflag.FlagSet) {
	flags.Bool("debug", false, "print CLI debugging information")
	flags.Lookup("debug").Hidden = true
}

func addMetaRepoFlag(flags *pflag.FlagSet) error {
	if homeDir, homeErr := os.UserHomeDir(); homeErr != nil {
		return fmt.Errorf("failed to locate home directory; %w", homeErr)
	} else {
		flags.String("meta-repo", filepath.Join(homeDir, "meta"), "Meta repo to use")
		return nil
	}
}

/* Flag configuration */

// Flag configuration for a single CLI command
type CommandFlagSet interface {
	//Add the implemented flags to the given CLI command
	AddTo(cmd *cobra.Command) error
}

/* Use */

// Parse application configuration from flags passed to the CLI
func ParseFlags(flags *pflag.FlagSet) (AppConfig, error) {
	if debug, debugErr := flags.GetBool("debug"); debugErr != nil {
		return nil, debugErr
	} else if metaRepoPath, metaRepoPathErr := flags.GetString("meta-repo"); metaRepoPathErr != nil {
		return nil, metaRepoPathErr
	} else {
		config := &FlagAppConfig{
			appFactory:   use.NewAppFactory(),
			debug:        debug,
			flagSet:      flags,
			metaRepoPath: metaRepoPath,
		}
		return config, nil
	}
}

// Application configuration derived from flags passed to the CLI.
type FlagAppConfig struct {
	appFactory   *use.AppFactory
	debug        bool
	flagSet      *pflag.FlagSet
	metaRepoPath string
}

func (config FlagAppConfig) AppFactory() *use.AppFactory { return config.appFactory }
func (config FlagAppConfig) Debug() bool                 { return config.debug }
func (config FlagAppConfig) MetaRepoPath() string        { return config.metaRepoPath }
func (config FlagAppConfig) PrintDebug(writer io.Writer) {
	fmt.Fprintf(writer, "Flags:\n")

	debugFlag := config.flagSet.Lookup("debug")
	fmt.Fprintf(writer, "- debug [%v]: %v\n", debugFlag.DefValue, debugFlag.Value)

	metaRepoFlag := config.flagSet.Lookup("meta-repo")
	fmt.Fprintf(writer, "- meta-repo [%v]: %v\n", metaRepoFlag.DefValue, metaRepoFlag.Value)
}
