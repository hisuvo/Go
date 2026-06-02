package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	var err error

	conectStr := os.Getenv("DB_URL")

	DB, err = pgx.Connect(context.Background(), conectStr)

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected successfuly")
}