package model

import (
	"time"
)

// id?: string;
// code?: string;
// name?: string;
// description?: string;
// price?: number;
// quantity?: number;
// inventoryStatus?: string;
// category?: string;
// image?: string;
// rating?: number;

// options?: ProductOptions[];

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Code        string    `json:"code" gorm:"type:varchar(100)"`
	Name        string    `json:"name" gorm:"type:varchar(255); not null"`
	Description string    `json:"description" gorm:"type:varchar(255)"` // unique??
	Price       uint      `json:"price" gorm:"not null"`
	Quantity    uint      `json:"quantity"`
	Status      string    `json:"status" gorm:"type:varchar(100)"`
	Image       string    `json:"image"`
	UserCreated uint      `json:"userCreated" gorm:"not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	// UpdatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	// CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// type ProductBasic struct {
// 	ID        uint      `json:"id"`
// 	Role      string    `json:"role"`
// 	Username  string    `json:"username"`
// 	Name      string    `json:"name"`
// 	Email     string    `json:"email"`
// 	CreatedAt time.Time `json:"created_at"`
// }

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
