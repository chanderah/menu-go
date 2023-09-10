package model

import (
	"time"
)

type Table struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"index:table_ix1;unique;not null" binding:"required"`
	Barcode string `json:"barcode" gorm:"not null"`
	Status  *bool  `json:"status"`

	CreatedAt time.Time `json:"createdAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
}
