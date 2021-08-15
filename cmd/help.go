package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "help: get help with this todo list.",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("help time")
	},
}

func init() {
	RootCmd.AddCommand(helpCmd)
}
