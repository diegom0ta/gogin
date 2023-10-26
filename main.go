package main

import (
	"errors"
	"log"

	"github.com/diegom0ta/gogin/book"
	"github.com/diegom0ta/gogin/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	books  = "/books"
	bookId = "/books/:id"
)

var (
	ErrFailedToConnect = errors.New("failed to connect database")
	ErrMigrationFailed = errors.New("database migration failed")
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(ErrFailedToConnect)
	}

	err = db.AutoMigrate(&book.Book{})
	if err != nil {
		log.Println(ErrMigrationFailed)
	}

	h := handler.NewHandler(db)

	r := gin.New()

	r.GET(books, h.ListBooksHandler)
	r.POST(books, h.CreateBookHandler)
	r.DELETE(bookId, h.DeleteBookHandler)

	r.Run()
}
