package cmd

import (
	"fmt"
	"os"
	"strconv"

	"example.com/todo-cli/internal/storage/csv"
	"example.com/todo-cli/internal/task"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		handleComplete(args)
	},
	Args:  cobra.MinimumNArgs(1),
	Use:   "complete [taskID]",
	Short: "Mark a task as completed",
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

func handleComplete(args []string) {

	taskId, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid taskId, %s\n", args[0])
		return
	}

	tasks, err := csv.LoadTasks(dataFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s retrieving tasks from storage.Try again.\n", err)
		return
	}

	if err := task.CompleteTask(tasks, taskId); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s occured while updating the task\n", err)
		return
	}

	if err := csv.SaveTasks(dataFile, tasks); err != nil {
		fmt.Fprintf(os.Stderr, "Error : %s occured while saving the tasks. Try again.\n", err)
		return
	}

	fmt.Printf("Your task with taskID: %d has been marked as completed.\n", taskId)

}
