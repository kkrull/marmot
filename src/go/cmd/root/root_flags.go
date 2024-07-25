package cmdroot

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var globalFlags = []rootFlag{debugFlag, metaRepoFlag}

type rootFlag string

const (
	debugFlag    rootFlag = "debug"
	metaRepoFlag rootFlag = "meta-repo"
)

func (flag rootFlag) AddTo(flags *pflag.FlagSet) error {
	switch flag {
	case debugFlag:
		flags.Bool(flag.Id(), false, "print CLI debugging information")
		flags.Lookup(flag.Id()).Hidden = true
		return nil

	case metaRepoFlag:
		if homeDir, homeErr := os.UserHomeDir(); homeErr != nil {
			return fmt.Errorf("failed to locate home directory; %w", homeErr)
		} else {
			flags.String(flag.Id(), filepath.Join(homeDir, "meta"), "Meta repo to use")
			return nil
		}

	default:
		return errors.Join(errors.ErrUnsupported, fmt.Errorf("unknown flag: %s", flag.Id()))
	}
}

func (flag rootFlag) Find(flags *pflag.FlagSet) string {
	flagObj := flags.Lookup(flag.Id())
	if flagObj == nil {
		return ""
	} else {
		return flagObj.Value.String()
	}
}

func (flag rootFlag) Id() string {
	return string(flag)
}

func (flag rootFlag) LongName() string {
	switch flag {
	case debugFlag:
		return "debug"
	case metaRepoFlag:
		return "meta-repo"
	default:
		return fmt.Sprintf("unknown flag: %s", flag.Id())
	}
}

// Flag configuration for the root (e.g. top-level) command that dispatches to all other commands.
func RootFlagSet() CommandFlags {
	return &rootFlags{}
}

// Flags that can be passed to a CLI command.
type CommandFlags interface {
	// Add the implemented flags to the given CLI command.
	AddTo(cmd *cobra.Command) error
}

type rootFlags struct{}

func (rootFlags) AddTo(rootCmd *cobra.Command) error {
	debugFlag.AddTo(rootCmd.PersistentFlags())
	return metaRepoFlag.AddTo(rootCmd.PersistentFlags())
}
