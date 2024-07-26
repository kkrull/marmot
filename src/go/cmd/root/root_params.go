package cmdroot

import (
	"fmt"
	"io"
	"net/url"

	"github.com/kkrull/marmot/core"
	"github.com/kkrull/marmot/svcfs"
	"github.com/kkrull/marmot/use"
	"github.com/spf13/pflag"
)

// Parameters passed to the root command
func RootCommandParser() (CommandParser, error) {
	if version, versionErr := core.MarmotVersion(); versionErr != nil {
		return nil, versionErr
	} else {
		return &rootParamParser{version: version}, nil
	}
}

// Parses parameters passed to a CLI command through environment, flags, and positional arguments.
type CommandParser interface {
	Parse(flags *pflag.FlagSet, args []string) (AppConfig, error)
	ParseR(flags *pflag.FlagSet, args []string, stdin io.Reader) (AppConfig, error)
}

type rootParamParser struct {
	version string
}

func (parser rootParamParser) Parse(flags *pflag.FlagSet, args []string) (AppConfig, error) {
	if debug, debugErr := debugFlag.GetBool(flags); debugErr != nil {
		return nil, debugErr
	} else if metaRepoPath, pathErr := metaRepoFlag.GetString(flags); pathErr != nil {
		return nil, pathErr
	} else {
		config := &rootParams{
			appFactory: use.NewAppFactory().
				WithMetaDataAdmin(svcfs.NewJsonMetaRepoAdmin(parser.version)).
				WithRepositorySource(svcfs.NewJsonMetaRepo(metaRepoPath)),
			args:         args,
			debug:        debug,
			flagSet:      flags,
			metaRepoPath: metaRepoPath,
		}

		return config, nil
	}
}

func (parser rootParamParser) ParseR(flags *pflag.FlagSet, args []string, stdin io.Reader) (AppConfig, error) {
	if debug, debugErr := debugFlag.GetBool(flags); debugErr != nil {
		return nil, debugErr
	} else if metaRepoPath, pathErr := metaRepoFlag.GetString(flags); pathErr != nil {
		return nil, pathErr
	} else {
		config := &rootParams{
			appFactory: use.NewAppFactory().
				WithMetaDataAdmin(svcfs.NewJsonMetaRepoAdmin(parser.version)).
				WithRepositorySource(svcfs.NewJsonMetaRepo(metaRepoPath)),
			args:         args,
			debug:        debug,
			flagSet:      flags,
			metaRepoPath: metaRepoPath,
		}

		return config, nil
	}
}

// Application configuration derived from flags passed to the CLI.
type rootParams struct {
	appFactory   use.AppFactory
	args         []string
	debug        bool
	flagSet      *pflag.FlagSet
	metaRepoPath string
}

func (params rootParams) AppFactory() use.AppFactory { return params.appFactory }
func (params rootParams) Args() []string             { return params.args }

func (params rootParams) ArgsAsUrls() ([]*url.URL, error) {
	urls := make([]*url.URL, len(params.args))
	for i, rawArg := range params.args {
		if urlArg, parseErr := url.Parse(rawArg); parseErr != nil {
			return nil, fmt.Errorf("url expected: %s; %w", rawArg, parseErr)
		} else {
			urls[i] = urlArg
		}
	}

	return urls, nil
}

func (params rootParams) Debug() bool          { return params.debug }
func (params rootParams) MetaRepoPath() string { return params.metaRepoPath }
func (params rootParams) PrintDebug(writer io.Writer) {
	for i, arg := range params.args {
		fmt.Fprintf(writer, "[%d]: %s\n", i, arg)
	}

	for _, flag := range rootFlags {
		fmt.Fprintf(writer, "--%s=%s\n", flag.LongName(), flag.Find(params.flagSet))
	}
}
