package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() *PostgresStorage {
	connStr := "host=localhost port=5436 user=postgres password=mypass dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return &PostgresStorage{db: db}
}
