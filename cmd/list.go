package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var listCmd = &cobra.Command{
	Use: "list",	
	Short: "list out your todos",

	Run: func(cmd *cobra.Command, args []string ) {
		fmt.Println("listing out your todos")
	},
}

func init() { 
	RootCmd.AddCommand(listCmd)	
}