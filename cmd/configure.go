package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func NewConfigureCmd() *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			prompt := promptui.Select{
				Label: "What storage you want to save your tasks",
				Items: []string{"csv_file", "postgres"},
			}

			_, val, err := prompt.Run()

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			mpp, err := godotenv.Read()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			mpp["STORAGE_TYPE"] = val

			if err := godotenv.Write(mpp, ".env"); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			fmt.Printf("Storage type changed to %s", val)
		},
		Use:   "configure",
		Short: "configure the storage location of tasks",
	}
}
