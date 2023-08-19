package model

import (
	"time"
)

type ProductOption struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ProductId uint      `json:"productId" gorm:"not null;index:product_option_ix1" binding:"required"`
	Name      string    `json:"name,omitempty" gorm:"type:varchar(100)"`
	Option    string    `json:"option,omitempty"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
}

type Option struct {
	Value string `json:"value"`
	Price uint   `json:"price"`
}
