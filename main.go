package main

import (
	"context"
	"fmt"
	"os"

	"example.com/todo-cli/cmd"
	"example.com/todo-cli/internal/storage/csv"
	"example.com/todo-cli/internal/storage/postgres"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	ctx := context.Background()

	storageType := os.Getenv("STORAGE_TYPE")

	switch storageType {
	case "postgres":
		dbUrl := os.Getenv("POSTGRES_DB_URL")
		db, err := postgres.NewPostgresStorage(ctx, dbUrl)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		defer db.Conn.Close(ctx)
		cmd.Execute(db)

	case "csv_file":
		dbUrl := os.Getenv("CSV_FILE_PATH")
		db, err := csv.NewCSVStorage(dbUrl)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		cmd.Execute(db)
	}

}
