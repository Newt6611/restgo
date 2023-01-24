package main

import (
	"database/sql"
	"fmt"
	"os"
	"restgo/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	initHelper()
	router := gin.Default()
	router.GET("/", handler.RootHandler)
	router.GET("/ping", handler.RootHandler)

	v1 := router.Group("/v1")
	v1.GET("/books", handler.GetAllBooks)
	v1.GET("/books/:id", handler.GetOneBook)
	v1.POST("/books", handler.CreateOneBook)
	v1.PUT("/books", handler.UpdateBook)
	v1.DELETE("/books/:id", handler.DeleteOneBook)

	router.Run(":8081")
}

func initHelper() {
	// init db
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	handler.Init(db)
}
