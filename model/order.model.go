package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID     `json:"id" gorm:"primaryKey"`
	TableId    uint          `json:"tableId" gorm:"index:order_ix1;not null" binding:"required"`
	TotalPrice uint          `json:"totalPrice" gorm:"not null" binding:"required"`
	Status     *bool         `json:"status,omitempty"`
	Products   ProductsBasic `json:"products" gorm:"type:json;not null" binding:"required"`

	CreatedAt time.Time `json:"createdAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
}

func (o *Order) GenerateUUID() error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	o.ID = uuid
	return nil
}

type ProductsBasic []ProductBasic

type ProductBasic struct {
	ID       uint         `json:"id" gorm:"primaryKey"`
	Name     string       `json:"name" gorm:"type:varchar(255);not null;index:product_ix1;index:product_ix2" binding:"required" `
	Price    uint         `json:"price" gorm:"not null;index:product_ix1;index:product_ix2" binding:"required"`
	Image    string       `json:"image,omitempty"`
	Options  OptionsBasic `json:"options" gorm:"type:json"`
	Quantity uint         `json:"quantity,omitempty"`
}

type OptionsBasic []OptionBasic

type OptionBasic struct {
	Name   string             `json:"name"`
	Values []OptionValueBasic `json:"values"`
}

type OptionValueBasic struct {
	Value string `json:"value"`
	Price uint   `json:"price"`
}

func (o *ProductsBasic) Scan(value interface{}) error {
	b, _ := value.([]byte)
	return json.Unmarshal(b, &o)
}

func (o ProductsBasic) Value() (driver.Value, error) {
	return json.Marshal(&o)
}
