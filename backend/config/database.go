package config

import (
	"database/sql"
	"log"
)

func InitDB() *sql.DB {
	// Connection string
	connStr := "postgres://postgres:postgres@localhost:5432/godock?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
