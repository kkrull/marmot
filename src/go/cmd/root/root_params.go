package cmdroot

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"strings"

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
		metaRepoAdmin := svcfs.NewJsonMetaRepoAdmin(parser.version)
		jsonMetaRepo := svcfs.NewJsonMetaRepo(metaRepoPath)
		config := &rootParams{
			appFactory: use.NewAppFactory().
				WithMetaDataAdmin(metaRepoAdmin).
				WithRepositorySource(jsonMetaRepo),
			args:         args,
			debug:        debug,
			flagSet:      flags,
			inputLines:   make([]string, 0),
			metaRepoPath: metaRepoPath,
			queryFactory: use.NewQueryFactory().
				WithMetaDataAdmin(metaRepoAdmin).
				WithRepositorySource(jsonMetaRepo),
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
		argsBeforeDash := make([]string, 0)
		for _, arg := range args {
			if strings.TrimSpace(arg) == "-" {
				break
			} else {
				argsBeforeDash = append(argsBeforeDash, arg)
			}
		}

		inputLines := make([]string, 0)
		scanner := bufio.NewScanner(stdin)
		for scanner.Scan() {
			line := scanner.Text()
			inputLines = append(inputLines, line)
		}
		if scanErr := scanner.Err(); scanErr != nil {
			return nil, scanErr
		}

		metaRepoAdmin := svcfs.NewJsonMetaRepoAdmin(parser.version)
		jsonMetaRepo := svcfs.NewJsonMetaRepo(metaRepoPath)
		config := &rootParams{
			appFactory: use.NewAppFactory().
				WithMetaDataAdmin(metaRepoAdmin).
				WithRepositorySource(jsonMetaRepo),
			args:         argsBeforeDash,
			debug:        debug,
			flagSet:      flags,
			inputLines:   inputLines,
			metaRepoPath: metaRepoPath,
			queryFactory: use.NewQueryFactory().
				WithMetaDataAdmin(metaRepoAdmin).
				WithRepositorySource(jsonMetaRepo),
		}

		return config, nil
	}
}

// Application configuration derived from flags passed to the CLI.
type rootParams struct {
	//Application interface
	appFactory   use.AppFactory
	queryFactory use.QueryFactory

	//CLI arguments
	args []string

	//CLI flags
	debug        bool
	flagSet      *pflag.FlagSet
	metaRepoPath string

	//CLI input
	inputLines []string
}

/* Application interface */

func (params rootParams) CommandFactory() use.AppFactory { return params.appFactory }
func (params rootParams) QueryFactory() use.QueryFactory { return params.queryFactory }

/* CLI arguments */

func (params rootParams) Args() []string { return params.args }
func (params rootParams) ArgsAsUrls() ([]*url.URL, error) {
	urls := make([]*url.URL, len(params.args))
	for i, rawArg := range params.args {
		if urlArg, parseErr := url.Parse(rawArg); parseErr != nil {
			return nil, fmt.Errorf("url expected: <%s>; %w", rawArg, parseErr)
		} else {
			urls[i] = urlArg
		}
	}

	return urls, nil
}

/* CLI debugging */

func (params rootParams) Debug() bool { return params.debug }
func (params rootParams) PrintDebug(writer io.Writer) {
	for i, arg := range params.args {
		fmt.Fprintf(writer, "arg [%d]: %s\n", i, arg)
	}

	for _, flag := range rootFlags {
		fmt.Fprintf(writer, "flag --%s=%s\n", flag.LongName(), flag.Find(params.flagSet))
	}

	for i, line := range params.inputLines {
		fmt.Fprintf(writer, "stdin [%d]: %s\n", i, line)
	}
}

/* CLI flags */

func (params rootParams) MetaRepoPath() string { return params.metaRepoPath }

/* CLI input */

func (params rootParams) InputLines() []string { return params.inputLines }
func (params rootParams) InputLinesAsUrls() ([]*url.URL, error) {
	urls := make([]*url.URL, 0)
	for _, rawLine := range params.inputLines {
		trimmedLine := strings.TrimSpace(rawLine)
		if trimmedLine == "" {
			continue
		} else if urlLine, parseErr := url.Parse(rawLine); parseErr != nil {
			return nil, fmt.Errorf("url expected: <%s>; %w", rawLine, parseErr)
		} else {
			urls = append(urls, urlLine)
		}
	}

	return urls, nil
}
