package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kkrull/marmot/use"
	"github.com/spf13/cobra"
)

func ParseFlags(cobraCmd *cobra.Command) *Config {
	//TODO KDK: Look up flags and apply default configuration, for meta repo path
	return &Config{
		AppFactory: *use.NewAppFactory(),
	}
}

type Config struct {
	AppFactory use.AppFactory
}

func (config Config) MetaRepoPath() (string, error) {
	return defaultMetaRepoPath()
}

func defaultMetaRepoPath() (string, error) {
	if homeDir, homeErr := os.UserHomeDir(); homeErr != nil {
		return "", fmt.Errorf("failed to locate home directory; %w", homeErr)
	} else {
		return filepath.Join(homeDir, "meta"), nil
	}
}
