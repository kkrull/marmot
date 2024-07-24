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

/* Configuration */

func AddFlags(cobraCmd *cobra.Command) error {
	addDebugFlag(cobraCmd)
	if metaRepoErr := addMetaRepoFlag(cobraCmd); metaRepoErr != nil {
		return metaRepoErr
	} else {
		return nil
	}
}

func addDebugFlag(cobraCmd *cobra.Command) {
	cobraCmd.PersistentFlags().Bool("debug", false, "print CLI debugging information")
	cobraCmd.PersistentFlags().Lookup("debug").Hidden = true
}

func addMetaRepoFlag(cobraCmd *cobra.Command) error {
	if homeDir, homeErr := os.UserHomeDir(); homeErr != nil {
		return fmt.Errorf("failed to locate home directory; %w", homeErr)
	} else {
		cobraCmd.PersistentFlags().String(
			"meta-repo",
			filepath.Join(homeDir, "meta"),
			"Meta repo to use",
		)
		return nil
	}
}

/* Use */

func ParseFlags(cobraCmd *cobra.Command) (AppConfig, error) {
	flags := cobraCmd.Flags()
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
