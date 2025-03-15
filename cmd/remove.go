package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [task_id]",
	Short: "Remove a task by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err == nil {
			main.removeTask(id)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
