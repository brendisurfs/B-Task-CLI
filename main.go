package main

import (
	"github.com/brendisurfs/go-cli/cmd"
)

func main() {

	// runs our root command if no other args are passed into it.
	cmd.RootCmd.Execute()
}