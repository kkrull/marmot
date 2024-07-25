package cmdroot

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
)

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
