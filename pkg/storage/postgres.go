package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() *PostgresStorage {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5436"
	}
	connStr := fmt.Sprintf("host=%s port=%s user=postgres password=mypass dbname=test sslmode=disable", host, port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return &PostgresStorage{db: db}
}
