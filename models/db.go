package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func acquireDBConn() *sql.DB {
	db, err := sql.Open("sqlite3", "books.db")
	if err != nil {
		panic(err)
	}
	return db
}
