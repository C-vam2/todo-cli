package cmd

import (
	"fmt"
	"os"

	"example.com/todo-cli/internal/task"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("This is when the len of args is zero")
		} else {
			fmt.Println("This is with nonzero args length")
		}
	},
	Use:   "todo-cli",
	Short: "A CLI todo manager",
}

func Execute(repo task.TaskRepository) {

	rootCmd.AddCommand(
		NewAddCmd(repo),
		NewDeleteCmd(repo),
		NewListCmd(repo),
		NewUpdateCmd(repo),
		NewConfigureCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
