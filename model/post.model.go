package model

import (
	"time"
)

type Post struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}