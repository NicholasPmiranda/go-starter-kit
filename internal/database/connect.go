package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib" // Driver pgx para database/sql
	"github.com/jmoiron/sqlx"
)

// SQLXAdapter é um adaptador que permite usar sqlx.DB com a interface DBTX
type SQLXAdapter struct {
	DB *sqlx.DB
}

// getConnectionString retorna a string de conexão para o banco de dados PostgreSQL
func getConnectionString() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
}

// ConnectDB conecta ao banco de dados usando pgx e retorna uma conexão pgx
func ConnectDB() (*pgx.Conn, context.Context) {
	connString := getConnectionString()

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	log.Println("Successfully connected to database using pgx")
	ctx := context.Background()
	return conn, ctx
}

// ConnectDBX conecta ao banco de dados usando sqlx e retorna um DB sqlx
func ConnectDBX() *sqlx.DB {
	connString := getConnectionString()

	// Usando o driver pgx com sqlx
	db, err := sqlx.Connect("pgx", connString)
	if err != nil {
		log.Fatalf("Unable to connect to database with sqlx: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping database with sqlx: %v", err)
	}

	log.Println("Successfully connected to database using sqlx")
	return db
}
