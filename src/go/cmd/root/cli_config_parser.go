package cmdroot

import (
	"bufio"
	"io"
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

// TODO KDK: Why is this called from so many places?  Wouldn't it be better to pass CliConfig around to other commands?
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
		return parser.assemble(args, debug, flags, make([]string, 0), metaRepoPath), nil
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
		readFromStdin := false
		for _, arg := range args {
			if strings.TrimSpace(arg) == "-" {
				readFromStdin = true
				break
			} else {
				argsBeforeDash = append(argsBeforeDash, arg)
			}
		}

		if !readFromStdin {
			return parser.assemble(argsBeforeDash, debug, flags, make([]string, 0), metaRepoPath), nil
		} else if inputLines, scanErr := readLines(stdin); scanErr != nil {
			return nil, scanErr
		} else {
			return parser.assemble(argsBeforeDash, debug, flags, inputLines, metaRepoPath), nil
		}
	}
}

func (parser rootConfigParser) assemble(
	args []string,
	debug bool,
	flags *pflag.FlagSet,
	inputLines []string,
	metaRepoPath string,
) CliConfig {
	metaRepoAdmin := svcfs.NewJsonMetaRepoAdmin(parser.version)
	jsonMetaRepo := svcfs.NewJsonMetaRepo(metaRepoPath)
	return &rootCliConfig{
		args: args,
		cmdFactory: use.NewCommandFactory().
			WithLocalRepositorySource(jsonMetaRepo).
			WithMetaDataAdmin(metaRepoAdmin).
			WithRemoteRepositorySource(jsonMetaRepo),
		debug:        debug,
		flagSet:      flags,
		inputLines:   inputLines,
		metaRepoPath: metaRepoPath,
		queryFactory: use.NewQueryFactory().
			WithLocalRepositorySource(jsonMetaRepo).
			WithMetaDataAdmin(metaRepoAdmin).
			WithRemoteRepositorySource(jsonMetaRepo),
	}
}

func readLines(stdin io.Reader) ([]string, error) {
	inputLines := make([]string, 0)
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()
		inputLines = append(inputLines, line)
	}

	if scanErr := scanner.Err(); scanErr != nil {
		return nil, scanErr
	} else {
		return inputLines, nil
	}
}
