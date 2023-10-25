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

var ErrFailedToConnect = errors.New("failed to connect database")

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(ErrFailedToConnect)
	}

	db.AutoMigrate(&book.Book{})

	handler := handler.NewHandler(db)

	r := gin.New()

	r.GET("/books", handler.ListBooksHandler)
	r.POST("/books", handler.CreateBookHandler)
	r.DELETE("/books/:id", handler.DeleteBookHandler)

	r.Run()
}
