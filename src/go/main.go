package main

import (
	"os"

	"github.com/kkrull/marmot/cmd"
)

func main() {
	if executeErr := cmd.Execute(); executeErr != nil {
		os.Exit(1)
	}
}
