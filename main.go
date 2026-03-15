package main

import (
	"context"
	"fmt"
	"os"

	"example.com/todo-cli/cmd"
	"example.com/todo-cli/internal/storage/postgres"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	ctx := context.Background()
	dbUrl := os.Getenv("DB_URL")

	db, err := postgres.NewPostgresStorage(ctx, dbUrl)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer db.Conn.Close(ctx)
	// var repo task.TaskRepository
	cmd.Execute(db)

}
