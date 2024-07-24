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
	cobraCmd.PersistentFlags().Bool("debug", false, "print CLI debugging information")
	cobraCmd.PersistentFlags().Lookup("debug").Hidden = true
	if defaultPath, pathErr := defaultMetaRepoPath(); pathErr != nil {
		return pathErr
	} else {
		cobraCmd.PersistentFlags().String("meta-repo", defaultPath, "Meta repo to use")
		return nil
	}
}

func defaultMetaRepoPath() (string, error) {
	if homeDir, homeErr := os.UserHomeDir(); homeErr != nil {
		return "", fmt.Errorf("failed to locate home directory; %w", homeErr)
	} else {
		return filepath.Join(homeDir, "meta"), nil
	}
}

/* Use */

func ParseFlags(cobraCmd *cobra.Command) (*Config, error) {
	flags := cobraCmd.Flags()
	if debug, debugErr := flags.GetBool("debug"); debugErr != nil {
		return nil, debugErr
	} else if metaRepoPath, metaRepoPathErr := flags.GetString("meta-repo"); metaRepoPathErr != nil {
		return nil, metaRepoPathErr
	} else {
		return &Config{
			AppFactory:   *use.NewAppFactory(),
			Debug:        debug,
			MetaRepoPath: metaRepoPath,
			flagSet:      flags,
		}, nil
	}
}

type Config struct {
	AppFactory   use.AppFactory
	Debug        bool
	MetaRepoPath string
	flagSet      *pflag.FlagSet
}

func (config Config) PrintDebug(writer io.Writer) {
	fmt.Fprintf(writer, "Flags:\n")

	debugFlag := config.flagSet.Lookup("debug")
	fmt.Fprintf(writer, "- debug [%v]: %v\n", debugFlag.DefValue, debugFlag.Value)

	metaRepoFlag := config.flagSet.Lookup("meta-repo")
	fmt.Fprintf(writer, "- meta-repo [%v]: %v\n", metaRepoFlag.DefValue, metaRepoFlag.Value)
}
