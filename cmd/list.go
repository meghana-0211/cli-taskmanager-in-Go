package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		main.listTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
