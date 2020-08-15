package models

import "log"

type Author struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

func GetAuthor(author_id uint32) Author {
	db := acquireDBConn()
	defer db.Close()

	rows, err := db.Query("select name from author where id = ?", author_id)
	if err != nil {
		log.Print(err)
		return Author{}
	}

	var name string
	for rows.Next() {
		rows.Scan(&name)
	}
	return Author{Id: author_id, Name: name}
}
