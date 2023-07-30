package model

import "time"

type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name string `json:"name"`
	Role string `json:"role"`
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