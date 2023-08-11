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
	// ID          uint      `json:"id" gorm:"primaryKey"`
	// Code        string    `json:"code,omitempty" gorm:"type:varchar(100);index:product_ix1;index:product_ix2"`
	// Category    string    `json:"category,omitempty" gorm:"type:varchar(100);index:product_ix2"`
	// Name        string    `json:"name" gorm:"type:varchar(255);not null;index:product_ix1;index:product_ix2" binding:"required" `
	// Description string    `json:"description,omitempty" gorm:"type:varchar(255)"`
	// Price       uint      `json:"price" gorm:"not null;index:product_ix1;index:product_ix2" binding:"required"`
	// Quantity    uint      `json:"quantity,omitempty"`
	// Status      string    `json:"status,omitempty" gorm:"type:varchar(100)"`
	// Image       string    `json:"image,omitempty"`
	// UserCreated uint      `json:"userCreated,omitempty" gorm:"not null" binding:"required" `
	// UpdatedAt   time.Time `json:"updatedAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	// CreatedAt   time.Time `json:"createdAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	// CountRows   int64     `json:"countRows"`

	ID          uint      `json:"id" gorm:"primaryKey"`
	Code        string    `json:"code,omitempty" gorm:"type:varchar(100)"`
	Category    string    `json:"category,omitempty" gorm:"type:varchar(100)"`
	Name        string    `json:"name" gorm:"type:varchar(255); not null" binding:"required" `
	Description string    `json:"description,omitempty" gorm:"type:varchar(255)"`
	Price       uint      `json:"price" gorm:"not null" binding:"required"`
	Quantity    uint      `json:"quantity,omitempty"`
	Status      string    `json:"status,omitempty" gorm:"type:varchar(100)"`
	Image       string    `json:"image,omitempty"`
	UserCreated uint      `json:"userCreated,omitempty" gorm:"not null" binding:"required" `
	UpdatedAt   time.Time `json:"updatedAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	CreatedAt   time.Time `json:"createdAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	// CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
