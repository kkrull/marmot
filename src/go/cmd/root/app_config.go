package cmdroot

import (
	"io"
	"net/url"

	"github.com/kkrull/marmot/use"
)

// Abstract factory that creates various application factories from a CLI configuration.
type AppConfig interface {
	/* Application interface */

	//Constructs application commands.
	CommandFactory() use.CommandFactory

	//Constructs application queries.
	QueryFactory() use.QueryFactory

	/* CLI arguments */

	//Positional arguments that weren't parsed as flags, as raw text.
	Args() []string

	//Positional arguments parsed as URLs.
	ArgsAsUrls() ([]*url.URL, error)

	/* CLI debugging */

	//Print information for debugging CLI parsing to the given writer.
	PrintDebug(writer io.Writer)

	/* CLI flags */

	//Persistent flag to show debugging information instead of parsing and executing commands.
	Debug() bool

	//Persistent flag for the path to the meta repo to use in any parsed commands or queries.
	MetaRepoPath() string

	/* CLI input */

	//Lines of text input from another process–e.g. through a pipe to standard input–as raw text.
	InputLines() []string

	//Lines of text input from another process–e.g. through a pipe to standard input–parsed as URLs.
	InputLinesAsUrls() ([]*url.URL, error)
}
