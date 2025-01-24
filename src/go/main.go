package main

import "github.com/kkrull/marmot/cmd"

func main() {
	cmd := cmd.NewRootCmd()
	cmd.Execute()
}
