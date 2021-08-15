package cmd

import (
	"fmt"
	"os"

	"github.com/TwinProduction/go-color"
	"github.com/brendisurfs/b-task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list out your todos",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetAllTasks()
		if err != nil {
			fmt.Println(color.Red+"something went wrong: "+color.Reset, err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println(color.Green + "you have no tasks to complete. Dont stress yourself out :)" + color.Reset)
			return
		}
		fmt.Println(color.Bold + "you have the following tasks:" + color.Reset)
		for i, task := range tasks {
			fmt.Printf(color.Blue+"%d. %s\n"+color.Reset, i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
