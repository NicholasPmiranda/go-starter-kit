package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() (*pgx.Conn, context.Context) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	log.Println("Successfully connected to database")
	ctx := context.Background()
	return conn, ctx
}
