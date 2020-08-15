package models

import "log"

type Book struct {
	Id     uint32 `json:"id"`
	Name   string `json:"name"`
	Author Author `json:"author"`
	Status bool   `json:"status"`
}

func (b Book) Persist() {
	db := acquireDBConn()
	defer db.Close()

	s, err := db.Prepare("insert into book (Id, name, authord_Id, status) values (?, ?, ?, ?)")
	if err != nil {
		log.Print(err)
		return
	}
	s.Exec(b.Id, b.Name, b.Author.Id, b.Status)
}

func (b Book) Delete(Id int) error {
	db := acquireDBConn()
	defer db.Close()

	statement, err := db.Prepare("delete from book where Id = ?")
	if err != nil {
		log.Print(err)
		return err
	}
	statement.Exec(Id)
	return nil
}

func GetAllBooks() []Book {
	var books []Book
	db := acquireDBConn()
	defer db.Close()

	rows, err := db.Query("select * from book")
	if err != nil {
		log.Print(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var (
			Id     int
			name   string
			author int
			status bool
		)
		rows.Scan(&Id, &name, &author, &status)
		books = append(books, Book{Id: uint32(Id), Name: name, Author: GetAuthor(uint32((author))), Status: status})
	}
	return books
}

func GetBook(id int) Book {
	db := acquireDBConn()
	defer db.Close()

	rows, err := db.Query("select * from book where id = ?", id)
	defer rows.Close()
	if err != nil {
		log.Print(err)
		return Book{}
	}
	var book Book
	var authorId uint32
	for rows.Next() {
		rows.Scan(&book.Id, &book.Name, &authorId, &book.Status)
	}
	book.Author = GetAuthor(authorId)
	return book
}
