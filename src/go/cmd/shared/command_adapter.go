package cmdshared

import (
	"io"

	"github.com/spf13/cobra"
)

type CobraCommandRunFn = func(cli *cobra.Command, args []string) error

func CobraCommandAdapter[TConfig any](
	parseConfig func(*cobra.Command, []string) (TConfig, error),
	useConfig func(config TConfig, stdout io.Writer) error,
) CobraCommandRunFn {
	return func(cli *cobra.Command, args []string) error {
		if config, parseErr := parseConfig(cli, args); parseErr != nil {
			return parseErr
		} else {
			return useConfig(config, cli.OutOrStdout())
		}
	}
}
