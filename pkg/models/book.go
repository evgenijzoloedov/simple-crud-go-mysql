package models

import (
	"github.com/evgenijzoloedov/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name   string `gorm:""json:"name"`
	Author string `json:"author"`
	Title  string `json:"title"`
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

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var gettingBook Book
	db := db.Where("ID=?", id).Find(&gettingBook)
	return &gettingBook, db
}

func DeleteBooks(id int64) (deletedBook Book) {
	db.Where("ID=?", id).Delete(&deletedBook)
	return
}
