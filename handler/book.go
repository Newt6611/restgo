package handler

import (
	"net/http"
	"restgo/model"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM books")

	if e := checkError(c, err, http.StatusNotFound); e {
		return
	}

	var books []model.Book
	for rows.Next() {
		var book model.Book
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.CreatedTime)
		if e := checkError(c, err, http.StatusNotFound); e {
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func GetOneBook(c *gin.Context) {
	id := c.Param("id")
	row := db.QueryRow("SELECT * FROM books WHERE id = $1", id)

	var book model.Book
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.CreatedTime)
	if e := checkError(c, err, http.StatusNotFound); e {
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateOneBook(c *gin.Context) {
	var input struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	err := c.Bind(&input)
	if e := checkError(c, err, http.StatusBadRequest); e {
		return
	}

	if input.Title == "" || input.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "title or author can't be empty",
		})
		return
	}

	_, err = db.Exec("INSERT INTO books(title, author, created_on) VALUES ($1, $2, $3);", input.Title, input.Author, time.Now().UTC())
	if e := checkError(c, err, http.StatusBadRequest); e {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":  input.Title,
		"author": input.Author,
	})
}

func UpdateBook(c *gin.Context) {
	var input struct {
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	err := c.Bind(&input)
	if e := checkError(c, err, http.StatusBadRequest); e {
		return
	}

	if input.Title == "" || input.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "title or author can't be empty",
		})
		return
	}

	_, err = db.Exec("UPDATE books SET title = $1, author = $2 WHERE id = $3", input.Title, input.Author, input.Id)
	if e := checkError(c, err, http.StatusBadRequest); e {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":  input.Title,
		"author": input.Author,
	})
}

func DeleteOneBook(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if e := checkError(c, err, http.StatusBadRequest); e {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
