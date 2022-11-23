package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	db *sql.DB
)

func Init(_db *sql.DB) {
	db = _db
}

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func checkError(c *gin.Context, err error, status int) bool {
	if err == nil {
		return false
	}
	c.AbortWithStatusJSON(status, gin.H{
		"error": err.Error(),
	})
	return true
}
