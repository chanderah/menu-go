package model

import (
	"time"
)

type Post struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}