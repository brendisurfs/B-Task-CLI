package cmd

import (
	"fmt"
	"os"

	"github.com/brendisurfs/go-cli/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list out your todos",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetAllTasks()
		if err != nil {
			fmt.Println("something went wrong: ", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("you have no tasks to complete. Dont stress yourself out :)")
			return
		}
		fmt.Println("you have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
