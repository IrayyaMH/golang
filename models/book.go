package models

import (
	"golang_project/pkg/config"

	_ "github.com/go-sql-driver/mysql"

	"database/sql"
)

var db *sql.DB

type Book struct {
	sql.Model
	Id          string `json:"id"`
	Name        string `sql:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.Prepare(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Query(" ")
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *sql.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(book)
	return book
}
