package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Role      string    `json:"role" gorm:"type:varchar(255)"`
	Username  string    `json:"username" gorm:"unique; type:varchar(100); not null"`
	Password  string    `json:"-" gorm:"not null"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255)"` // unique??
	CreatedAt time.Time `json:"created_at" gorm:"default:NOW()"`
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
