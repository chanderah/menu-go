package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey"`
	TableId    string    `json:"tableId" gorm:"index:order_ix1;unique;not null" binding:"required"`
	TotalPrice uint      `json:"totalPrice" gorm:"not null" binding:"required"`
	Products   Products  `json:"products" gorm:"type:json;not null" binding:"required"`
	Status     string    `json:"status,omitempty"`

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

type Products []Product

func (o *Products) Scan(value interface{}) error {
	b, _ := value.([]byte)
	return json.Unmarshal(b, &o)
}

func (o Products) Value() (driver.Value, error) {
	return json.Marshal(&o)
}
