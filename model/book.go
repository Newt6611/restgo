package model

import "time"

type Book struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	CreatedTime time.Time `json:"created_time"`
}

func (b Book) TableName() string {
	return "books"
}
