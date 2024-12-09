package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	var err error
	DB, err = pgx.Connect(context.Background(), "postgres://postgres:admin@localhost:5432/productdb")
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
}
