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

	// Constructs application commands.
	CommandFactory() use.CommandFactory

	// Constructs application queries.
	QueryFactory() use.QueryFactory

	/* CLI arguments */

	// Positional arguments that weren't parsed as flags, as raw text.
	Args() []string

	// Positional arguments parsed as URLs.
	ArgsAsUrls() ([]*url.URL, error)

	/* CLI debugging */

	// Print information for debugging CLI parsing to the given writer.
	PrintDebug(writer io.Writer)

	/* CLI flags */

	// Persistent flag to show debugging information instead of parsing and executing commands.
	Debug() bool

	// Persistent flag for the path to the meta repo to use in any parsed commands or queries.
	MetaRepoPath() string

	/* CLI input */

	// Lines of text input from another process–e.g. through a pipe to standard input–as raw text.
	InputLines() []string

	// Lines of text input from another process–e.g. through a pipe to standard input–parsed as URLs.
	InputLinesAsUrls() ([]*url.URL, error)
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
