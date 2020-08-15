package models

import (
	"log"
	"bufio"
	"database/sql"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func LoadData(args []string) {
	db, err := sql.Open("sqlite3", "books.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	username := args[0]
	password := args[1]
	filename := args[2]
	rows, err := db.Query("select role, password from user where username = ?", username)
	if err != nil {
		panic("Something went wrong while accessing db for authorization")
	}
	var origPassword string
	var role string
	for rows.Next() {
		rows.Scan(&role, &origPassword)
	}
	if origPassword != password || role != "admin" {
		panic("You are unauthorized to load the file")
	}

	file, err := os.Open(filename)
	if err != nil {
		panic("Error: unable to read file")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		values := strings.Split(row, ",")
		statement, err := db.Prepare("insert into book (id, name, author_id, status) values (?, ?, ?, ?)")
		if err != nil {
			log.Print(err)
			continue
		}
		statement.Exec(values[0], values[1], values[2], values[3])
	}
}

