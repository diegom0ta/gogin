package main

import (
	"errors"
	"log"

	"github.com/diegom0ta/gogin/book"
	"github.com/diegom0ta/gogin/handler"
	"github.com/diegom0ta/gogin/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	books  = "/books"
	bookId = "/books/:id"
	login  = "/login"
	root   = "/"
	testDb = "test.db"
)

var (
	ErrFailedToConnect = errors.New("failed to connect database")
	ErrMigrationFailed = errors.New("database migration failed")
)

func main() {
	db, err := gorm.Open(sqlite.Open(testDb), &gorm.Config{})
	if err != nil {
		log.Println(ErrFailedToConnect)
	}

	err = db.AutoMigrate(&user.User{}, &book.Book{})
	if err != nil {
		log.Println(ErrMigrationFailed)
	}

	h := handler.NewHandler(db)

	r := gin.New()

	r.POST(login, h.LoginHandler)

	protected := r.Group(root, h.AuthMid)

	protected.GET(books, h.ListBooksHandler)
	protected.POST(books, h.CreateBookHandler)
	protected.DELETE(bookId, h.DeleteBookHandler)

	r.Run()
}
