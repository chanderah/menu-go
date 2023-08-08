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
	Category    string    `json:"category" gorm:"type:varchar(100)"`
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

type ProductId struct {
	ID uint `json:"id" gorm:"binding:required"`
}

type CreateProduct struct {
	Name        string `json:"name" gorm:"binding:required"`
	Price       uint   `json:"price" gorm:"binding:required"`
	UserCreated uint   `json:"userCreated" gorm:"binding:required"`
}
