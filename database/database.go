package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_ "path/filepath"
)

var db *sql.DB

func GetDBClient() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	var err error

	db, err := sql.Open("sqlite3", "./database/database.db")

	if err != nil {
		return nil, err
	}

	return db, nil
}
