package models

import (
	"log"
)

const (
	Admin   = "admin"
	Regular = "regular"
)

type User struct {
	uid  uint32 `json:"uid"`
	name string `json:"name"`
	role string `json:"role"`
}

func Login(username, password string) (string, error) {
	db := acquireDBConn()
	defer db.Close()

	rows, err := db.Query("select password from user where username = ?", username)
	if err != nil {
		log.Fatal("Error occurred while fetching records from sqlite", err)
		return "", err
	}
	defer rows.Close()

	var origPassword string
	for rows.Next() {
		if err := rows.Scan(&origPassword); err != nil {
			log.Print(err)
		}
	}
	if password != origPassword {
		return "", nil
	}
	return "this is a token", nil
}
