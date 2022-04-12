package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
)

func main() {
	dsn := "phpmyadmin:inifgrup@tcp(103.16.133.234:3306)/khairon?charset=utf8mb4&parseTime=True&loc=Local"
	var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}
	fmt.Println("Sukses Konek ke DB")
	db.AutoMigrate(book.Book{})
	//===============
	//CREATE
	//===============
	bookRepository := book.NewRepository(db)
	book := book.Book{
		Title:       "Tangisan Ibu",
		Description: "Tangisan ibu untuk anaknya pada tuhannya",
		Price:       105000,
		Rating:      4,
		Discount:    0,
	}
	bookRepository.Create(book)
	//newBook, err := bookRepository.Create(book)
	//===============
	//FindByID
	//===============
	//book, err := bookRepository.FindByID(2)
	//fmt.Println("Title :", book.Title)
	//===============
	//FindAll
	//===============
	//bookRepository := book.NewRepository(db)
	//books, err := bookRepository.FindAll()
	//for _, book := range books {
	//	fmt.Println("Title :", book.Title)
	//}

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/about", handler.AboutHandler)
	v1.GET("/book/:id", handler.BookHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)
	router.Run()
}
