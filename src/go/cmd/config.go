package cmd

import "github.com/spf13/cobra"

/* Flags */

func DefaultGlobalFlags() *GlobalConfig {
	return &GlobalConfig{
		MetaRepoHome: "/home/me/meta-default",
	}
}

func ParseGlobalFlags(cmd *cobra.Command) *GlobalConfig {
	metaHomeFlag := cmd.Flags().Lookup("meta-home")
	return &GlobalConfig{
		MetaRepoHome: metaHomeFlag.Value.String(),
	}
}

type GlobalConfig struct {
	MetaRepoHome string
}
