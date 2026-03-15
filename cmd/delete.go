package cmd

import (
	"fmt"
	"os"
	"strconv"

	"example.com/todo-cli/internal/storage/csv"
	"example.com/todo-cli/internal/task"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		handleDelete(args)
	},
	Args:  cobra.MinimumNArgs(1),
	Use:   "delete [taskID]",
	Short: "Delete a given task",
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func handleDelete(args []string) {

	taskID, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid taskID, %s \n", args[0])
		return
	}

	tasks, err := csv.LoadTasks(dataFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	tasks, err = task.DeleteTask(tasks, taskID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	if err := csv.SaveTasks(dataFile, tasks); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	fmt.Printf("Task with taskID : %d deleted successfully\n", taskID)

}
