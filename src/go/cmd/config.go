package cmd

import (
	"github.com/kkrull/marmot/use"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func ParseFlags(cobraCmd *cobra.Command) *Config {
	return &Config{
		AppFactory: *use.NewAppFactory(),
		flagSet:    cobraCmd.Flags(),
	}
}

type Config struct {
	AppFactory use.AppFactory
	flagSet    *pflag.FlagSet
}

func (config Config) MetaRepoPath() string {
	return config.flagSet.Lookup("meta-repo").Value.String()
}
