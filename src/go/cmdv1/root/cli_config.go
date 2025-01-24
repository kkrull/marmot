package cmdroot

import (
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/kkrull/marmot/use"
	"github.com/spf13/pflag"
)

// Abstract factory that creates various application factories from a CLI configuration.
type CliConfig interface {
	/* Application interface */

	// Constructs application actions.
	ActionFactory() use.ActionFactory

	// Constructs application queries.
	QueryFactory() use.QueryFactory

	/* CLI arguments */

	// Positional arguments that weren't parsed as flags, as raw text.
	Args() []string

	// Positional arguments parsed as URLs.
	ArgsAsUrls() ([]*url.URL, error)

	// Positional arguments trimmed for whitespace.
	ArgsTrimmed() []string

	/* CLI debugging */

	// Print information for debugging CLI parsing to the given writer.
	PrintDebug(writer io.Writer)

	/* CLI flags */

	// Persistent flag to show debugging information instead of parsing and executing commands.
	Debug() bool

	// Persistent flag for the path to the meta repo to use in any parsed actions or queries.
	MetaRepoPath() string

	/* CLI input */

	// Lines of text input from another process–e.g. through a pipe to standard input–as raw text.
	InputLines() []string

	// Lines of text input from another process–e.g. read from standard input–parsed as URLs.
	InputLinesAsUrls() ([]*url.URL, error)

	// Lines of text input from another process–e.g. read from standard input–trimmed for whitespace.
	InputLinesTrimmed() []string
}

// Application configuration derived from flags passed to the CLI.
type rootCliConfig struct {
	// Application interface
	cmdFactory   use.ActionFactory
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

func (config rootCliConfig) ActionFactory() use.ActionFactory { return config.cmdFactory }
func (config rootCliConfig) QueryFactory() use.QueryFactory   { return config.queryFactory }

/* CLI arguments */

func (config rootCliConfig) Args() []string { return config.args }

func (config rootCliConfig) ArgsAsUrls() ([]*url.URL, error) {
	urls := make([]*url.URL, len(config.args))
	for i, rawArg := range config.args {
		if urlArg, parseErr := url.Parse(rawArg); parseErr != nil {
			return nil, fmt.Errorf("url expected: <%s>; %w", rawArg, parseErr)
		} else {
			urls[i] = urlArg
		}
	}

	return urls, nil
}

func (config rootCliConfig) ArgsTrimmed() []string {
	trimmed := make([]string, len(config.args))
	for i, rawArg := range config.args {
		trimmed[i] = strings.TrimSpace(rawArg)
	}

	return trimmed
}

/* CLI debugging */

func (config rootCliConfig) Debug() bool { return config.debug }

func (config rootCliConfig) PrintDebug(writer io.Writer) {
	for i, arg := range config.args {
		fmt.Fprintf(writer, "arg [%d]: %s\n", i, arg)
	}

	for _, flag := range rootFlags {
		fmt.Fprintf(writer, "flag --%s=%s\n", flag.LongName(), flag.Find(config.flagSet))
	}

	for i, line := range config.inputLines {
		fmt.Fprintf(writer, "stdin [%d]: %s\n", i, line)
	}
}

/* CLI flags */

func (config rootCliConfig) MetaRepoPath() string { return config.metaRepoPath }

/* CLI input */

func (config rootCliConfig) InputLines() []string { return config.inputLines }

func (config rootCliConfig) InputLinesAsUrls() ([]*url.URL, error) {
	urls := make([]*url.URL, 0)
	for _, rawLine := range config.inputLines {
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

func (config rootCliConfig) InputLinesTrimmed() []string {
	trimmed := make([]string, len(config.inputLines))
	for i, rawLine := range config.inputLines {
		trimmed[i] = strings.TrimSpace(rawLine)
	}

	return trimmed
}
