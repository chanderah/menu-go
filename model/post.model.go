package model

import (
	"time"
)

type Post struct {
	// gorm.Model
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"type: varchar(255); not null" binding:"required"`
	Content string `json:"content" gorm:"not null" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"default:NOW()"`
}