package model

import (
	"time"
)

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Role      string    `json:"role" gorm:"type:varchar(255)"`
	Username  string    `json:"username" gorm:"unique; type:varchar(100); not null"`
	Password  string    `json:"password" gorm:"not null"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255)"` // unique??
	Token     string    `json:"token" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	// UpdatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	// CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type UserBasic struct {
	Id        uint      `json:"id"`
	Role      string    `json:"role"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}
