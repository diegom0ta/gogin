package main

import (
	"github.com/diegom0ta/gogin/book"
	"github.com/diegom0ta/gogin/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&book.Book{})

	handler := handler.NewHandler(db)

	r := gin.New()

	r.GET("/books", handler.ListBooksHandler)
	r.POST("/books", handler.CreateBookHandler)
	r.DELETE("/books/:id", handler.DeleteBookHandler)

	r.Run()
}
