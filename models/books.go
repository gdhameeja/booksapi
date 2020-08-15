package models

type Book struct {
	id uint32 `json:"id"`
	name string `json:"name"`
	author Author `json:"author"`
	status bool `json:"status"`
}

