package infrastructure

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgresDB(dsn string) *sql.DB {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Default().Fatalf("failed to open database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failes to ping database: %v", err)
	}
	return db
}
