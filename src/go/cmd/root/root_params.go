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

// Parses configuration from the arguments, environment, flags, and input available to the CLI.
type CliConfigParser interface {
	Parse(flags *pflag.FlagSet, args []string) (CliConfig, error)
	ParseR(flags *pflag.FlagSet, args []string, stdin io.Reader) (CliConfig, error)
}

// Parses configuration from arguments, flags, and input to the root command.
func RootConfigParser() (CliConfigParser, error) {
	if version, versionErr := core.MarmotVersion(); versionErr != nil {
		return nil, versionErr
	} else {
		return &rootConfigParser{version: version}, nil
	}
}

// Parse configuration that applies to the root command and its descendant sub-commands.
type rootConfigParser struct {
	version string
}

func (parser rootConfigParser) Parse(
	flags *pflag.FlagSet, args []string,
) (CliConfig, error) {
	if debug, debugErr := debugFlag.GetBool(flags); debugErr != nil {
		return nil, debugErr
	} else if metaRepoPath, pathErr := metaRepoFlag.GetString(flags); pathErr != nil {
		return nil, pathErr
	} else {
		metaRepoAdmin := svcfs.NewJsonMetaRepoAdmin(parser.version)
		jsonMetaRepo := svcfs.NewJsonMetaRepo(metaRepoPath)
		config := &rootCliConfig{
			args: args,
			cmdFactory: use.NewCommandFactory().
				WithMetaDataAdmin(metaRepoAdmin).
				WithRepositorySource(jsonMetaRepo),
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

func (parser rootConfigParser) ParseR(
	flags *pflag.FlagSet, args []string, stdin io.Reader,
) (CliConfig, error) {
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
		config := &rootCliConfig{
			args: argsBeforeDash,
			cmdFactory: use.NewCommandFactory().
				WithMetaDataAdmin(metaRepoAdmin).
				WithRepositorySource(jsonMetaRepo),
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
type rootCliConfig struct {
	// Application interface
	cmdFactory   use.CommandFactory
	queryFactory use.QueryFactory

	// CLI arguments
	args []string

	// CLI flags
	debug        bool
	flagSet      *pflag.FlagSet
	metaRepoPath string

	// CLI input
	inputLines []string
}

/* Application interface */

func (params rootCliConfig) CommandFactory() use.CommandFactory { return params.cmdFactory }
func (params rootCliConfig) QueryFactory() use.QueryFactory     { return params.queryFactory }

/* CLI arguments */

func (params rootCliConfig) Args() []string { return params.args }

func (params rootCliConfig) ArgsAsUrls() ([]*url.URL, error) {
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

func (params rootCliConfig) Debug() bool { return params.debug }

func (params rootCliConfig) PrintDebug(writer io.Writer) {
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

func (params rootCliConfig) MetaRepoPath() string { return params.metaRepoPath }

/* CLI input */

func (params rootCliConfig) InputLines() []string { return params.inputLines }

func (params rootCliConfig) InputLinesAsUrls() ([]*url.URL, error) {
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
