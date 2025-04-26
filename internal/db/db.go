package db

import (
	"database/sql"
	"os"

	_ "github.com/sijms/go-ora/v2"
)

func Connect() (*sql.DB, error) {
	dsn := os.Getenv("ORACLE_DSN")
	db, err := sql.Open("oracle", dsn)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}
