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
func RootCliParams() (CommandParamParser, error) {
	if version, versionErr := core.MarmotVersion(); versionErr != nil {
		return nil, versionErr
	} else {
		return &rootParamParser{version: version}, nil
	}
}

// Parses parameters passed to a CLI command through environment, flags, and positional arguments.
type CommandParamParser interface {
	ParseAppConfig(flags *pflag.FlagSet, args []string) (AppConfig, error)
}

type rootParamParser struct {
	version string
}

func (parser rootParamParser) ParseAppConfig(flags *pflag.FlagSet, args []string) (AppConfig, error) {
	if debug, debugErr := flags.GetBool(debugFlag); debugErr != nil {
		return nil, debugErr
	} else if metaRepoPath, metaRepoPathErr := flags.GetString(metaRepoFlag); metaRepoPathErr != nil {
		return nil, metaRepoPathErr
	} else {
		config := &globalParams{
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
type globalParams struct {
	appFactory   use.AppFactory
	args         []string
	debug        bool
	flagSet      *pflag.FlagSet
	metaRepoPath string
}

func (params globalParams) AppFactory() use.AppFactory { return params.appFactory }
func (params globalParams) Args() []string             { return params.args }

func (params globalParams) ArgsAsUrls() ([]*url.URL, error) {
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

func (params globalParams) Debug() bool          { return params.debug }
func (params globalParams) MetaRepoPath() string { return params.metaRepoPath }
func (params globalParams) PrintDebug(writer io.Writer) {
	for i, arg := range params.args {
		fmt.Fprintf(writer, "[%d]: %s\n", i, arg)
	}

	debugFlag := params.flagSet.Lookup(debugFlag)
	fmt.Fprintf(writer, "--debug [%v]: %v\n", debugFlag.DefValue, debugFlag.Value)

	metaRepoFlag := params.flagSet.Lookup(metaRepoFlag)
	fmt.Fprintf(writer, "--meta-repo [%v]: %v\n", metaRepoFlag.DefValue, metaRepoFlag.Value)
}
