package models

import (
	"github.com/igdevansh09/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publications"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
