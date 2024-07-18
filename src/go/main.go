package main

import "github.com/kkrull/marmot/cliversion"

func main() {
	versionCmd := cliversion.VersionCommand()
	versionCmd.Execute()
}
