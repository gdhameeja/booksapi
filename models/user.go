package models

import (
	"log"
)

const (
	Admin   = "admin"
	Regular = "regular"
)

type User struct {
	Id  uint32 `json:"id"`
	Username string `json:"Username"`
	Role string `json:"role"`
}

func Login(userUsername, password string) (string, error) {
	db := acquireDBConn()
	defer db.Close()

	rows, err := db.Query("select token, password from user where username = ?", userUsername)
	if err != nil {
		log.Fatal("Error occurred while fetching records from sqlite", err)
		return "", err
	}
	defer rows.Close()

	var origPassword string
	var token string
	for rows.Next() {
		if err := rows.Scan(&token, &origPassword); err != nil {
			log.Print(err)
		}
	}
	if password != origPassword {
		return "", nil
	}
	return token, nil
}

func GetUserForToken(token string) *User {
	db := acquireDBConn()
	defer db.Close()
	rows, err := db.Query("select id, Username, role from user where token = ?", token)
	if err != nil {
		log.Print(err)
		return nil
	}

	var user User
	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Role)
	}
	if user.Username == "" {
		return nil
	}
	return &user
}
