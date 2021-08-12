package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var doCmd = &cobra.Command {
	Use: "do",
	Short: "Marks a task as complete/doing!",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do called, get it done!")
	},
}

// adds our do command.
func init() {
	RootCmd.AddCommand(doCmd)
}