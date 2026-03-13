package main

import (
	"example.com/todo-cli/cmd"
	"github.com/joho/godotenv"
)

func main() {
	cmd.Execute()
	godotenv.Load()
}
