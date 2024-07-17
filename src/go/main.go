package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	versionCmd := &cobra.Command{
		Long: "marmot version prints the version number and exits",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("marmot-cli 0.0.1")
		},
		Short:   "print the version number",
		Use:     "marmot version",
		Version: "0.0.1",
	}
	versionCmd.Execute()
}
