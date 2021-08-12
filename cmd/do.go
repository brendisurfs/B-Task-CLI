package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)


var doCmd = &cobra.Command {
	Use: "do",
	Short: "Marks a task as complete/doing!",

	Run: func(cmd *cobra.Command, args []string) {
		var ids []int 

		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("error: failed to parse argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Println(ids)
	},
}

// adds our do command.
func init() {
	RootCmd.AddCommand(doCmd)
}