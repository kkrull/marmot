package cmdroot

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Flag configuration for the root (e.g. top-level) command that dispatches to all other commands.
func FlagSet() *rootFlagSet {
	return &rootFlagSet{}
}

// Flags that can be passed to a CLI command.
type rootFlagSet struct{}

// Add the implemented flags to the given CLI command.
func (rootFlagSet) AddTo(rootCmd *cobra.Command) error {
	var errorAcc []error = make([]error, 0)
	for _, f := range rootFlags {
		errorAcc = append(errorAcc, f.AddTo(rootCmd.PersistentFlags()))
	}

	return errors.Join(errorAcc...)
}

/* Root Flag enum */

var rootFlags = []rootFlag{debugFlag, metaRepoFlag}

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

func (flag rootFlag) GetBool(flags *pflag.FlagSet) (bool, error) {
	return flags.GetBool(flag.Id())
}

func (flag rootFlag) GetString(flags *pflag.FlagSet) (string, error) {
	return flags.GetString(flag.Id())
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
