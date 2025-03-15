package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task_name]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addTask(args[0]) // ✅ Correct function call
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
