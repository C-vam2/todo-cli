package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const dataFile = "tasks.csv"

var rootCmd *cobra.Command = &cobra.Command{
	Use:   "todo-cli",
	Short: "A CLI todod manager",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
