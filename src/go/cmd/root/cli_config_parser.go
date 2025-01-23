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
		jsonMetaRepo := svcfs.NewJsonMetaRepo(metaRepoPath) // TODO KDK: Work here - add another version that works on a .local file
		config := &rootCliConfig{
			args: args,
			cmdFactory: use.NewCommandFactory().
				WithLocalRepositorySource(jsonMetaRepo).
				WithMetaDataAdmin(metaRepoAdmin).
				WithRemoteRepositorySource(jsonMetaRepo),
			debug:        debug,
			flagSet:      flags,
			inputLines:   make([]string, 0),
			metaRepoPath: metaRepoPath,
			queryFactory: use.NewQueryFactory().
				WithLocalRepositorySource(jsonMetaRepo).
				WithMetaDataAdmin(metaRepoAdmin).
				WithRemoteRepositorySource(jsonMetaRepo),
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
		readFromStdin := false
		for _, arg := range args {
			if strings.TrimSpace(arg) == "-" {
				readFromStdin = true
				break
			} else {
				argsBeforeDash = append(argsBeforeDash, arg)
			}
		}

		inputLines := make([]string, 0)
		if readFromStdin {
			scanner := bufio.NewScanner(stdin)
			for scanner.Scan() {
				line := scanner.Text()
				inputLines = append(inputLines, line)
			}
			if scanErr := scanner.Err(); scanErr != nil {
				return nil, scanErr
			}
		}

		metaRepoAdmin := svcfs.NewJsonMetaRepoAdmin(parser.version)
		jsonMetaRepo := svcfs.NewJsonMetaRepo(metaRepoPath)
		config := &rootCliConfig{
			args: argsBeforeDash,
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

		return config, nil
	}
}
