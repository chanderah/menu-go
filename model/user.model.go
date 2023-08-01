package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Role      string    `json:"role" gorm:"type:varchar(255)"`
	Username  string    `json:"username" gorm:"unique; type:varchar(100); not null"`
	Password  string    `json:"password" gorm:"not null"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255)"` // unique??
	Token     string    `json:"token" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:NOW()"`
}

type UserBasic struct {
	ID        uint      `json:"id"`
	Role      string    `json:"role"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// CREATE TABLE tb_user (
//     id SERIAL PRIMARY KEY,
//     username TEXT,
//     password TEXT,
//     name TEXT,
//     role TEXT,
//     created_at TIMESTAMP DEFAULT NOW()
// );

// type User struct {
// 	ID uint `gorm:"primary_key"`
// 	Name string `gorm:"type:varchar(255);not null"`
// 	Email string `gorm:"uniqueIndex;not null"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }
