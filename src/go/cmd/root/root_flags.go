package cmdroot

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kkrull/marmot/svcfs"
	"github.com/kkrull/marmot/use"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

/* CLI command flag configuration */

// Flag configuration for the root (e.g. top-level) command that dispatches to all other commands.
func RootFlagSet() CommandFlagSet {
	return &rootFlagSet{}
}

type rootFlagSet struct{}

/* App configuration */

func (rootFlagSet) ParseAppConfig(flags *pflag.FlagSet, args []string) (AppConfig, error) {
	if debug, debugErr := flags.GetBool("debug"); debugErr != nil {
		return nil, debugErr
	} else if metaRepoPath, metaRepoPathErr := flags.GetString("meta-repo"); metaRepoPathErr != nil {
		return nil, metaRepoPathErr
	} else {
		config := &FlagAppConfig{
			appFactory: use.NewAppFactory().
				WithMetaDataAdmin(svcfs.NewJsonMetaRepoAdmin()).
				WithRepositorySource(svcfs.NewJsonMetaRepo(metaRepoPath)),
			args:         args,
			debug:        debug,
			flagSet:      flags,
			metaRepoPath: metaRepoPath,
		}

		return config, nil
	}
}

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