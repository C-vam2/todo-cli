package cmd

import (
	"fmt"
	"os"
	"strings"

	"example.com/todo-cli/internal/storage"
	"example.com/todo-cli/internal/task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		handleAdd(args)
	},
	Args: cobra.MinimumNArgs(1),

	Use:   "add [task]",
	Short: "Add a new task",
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func handleAdd(args []string) {

	newTask := strings.Join(args, " ")

	if newTask == "" {
		fmt.Fprintln(os.Stderr, "Your task is empty")
		return
	}

	tasks, err := storage.LoadTasks(dataFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	newTasks := task.AddTask(tasks, newTask)
	err = storage.SaveTasks(dataFile, newTasks)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("Added task:", newTask)
}
