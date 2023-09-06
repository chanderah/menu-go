package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Code        string    `json:"code,omitempty" gorm:"type:varchar(100);index:product_ix1;index:product_ix2"`
	CategoryId  uint      `json:"categoryId,omitempty" gorm:"index:product_ix2"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null;index:product_ix1;index:product_ix2" binding:"required" `
	Description string    `json:"description,omitempty" gorm:"type:varchar(100)"`
	Options     Options   `json:"options" gorm:"type:json"`
	Price       uint      `json:"price" gorm:"not null;index:product_ix1;index:product_ix2" binding:"required"`
	Quantity    uint      `json:"quantity,omitempty"`
	Image       string    `json:"image,omitempty"`
	Featured    *bool     `json:"featured,omitempty"`
	Status      *bool     `json:"status,omitempty"`
	UserCreated uint      `json:"userCreated,omitempty" gorm:"not null" `
	UpdatedAt   time.Time `json:"updatedAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	CreatedAt   time.Time `json:"createdAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
}

type Options []ProductOptions

type ProductOptions struct {
	Name     string                `json:"name"`
	Multiple *bool                 `json:"multiple"`
	Required *bool                 `json:"required"`
	Values   []ProductOptionValues `json:"values"`
}

type ProductOptionValues struct {
	Value string `json:"value"`
	Price uint   `json:"price"`
}

func (o *Options) Scan(value interface{}) error {
	b, _ := value.([]byte)
	return json.Unmarshal(b, &o)
}

func (o Options) Value() (driver.Value, error) {
	return json.Marshal(&o)
}
