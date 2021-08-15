package cmd

import (
	"github.com/TwinProduction/go-color"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   color.Bold + "b-task" + color.Reset,
	Short: "A lightweight CLI todo list\n written by Brendan Prednis",
	// dont need a long.
	// no run
}
