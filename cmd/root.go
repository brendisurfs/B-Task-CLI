package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: "cli-task", 
	Short: "a lightweight CLI todo list", 
	// dont need a long.
	// no run
}