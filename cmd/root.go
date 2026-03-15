package cmd

import (
	"fmt"
	"os"

	"example.com/todo-cli/internal/task"
	"github.com/spf13/cobra"
)

const dataFile = "tasks.csv"

var rootCmd *cobra.Command = &cobra.Command{
	Use:   "todo-cli",
	Short: "A CLI todo manager",
}

func Execute(repo task.TaskRepository) {

	rootCmd.AddCommand(
		NewAddCmd(repo),
		NewDeleteCmd(repo),
		NewListCmd(repo),
		NewUpdateCmd(repo),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
