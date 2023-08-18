package model

import "time"

type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Order       int       `json:"order" gorm:"not null" binding:"required"`
	Name        string    `json:"name" gorm:"type:varchar(100);unique;not null" binding:"required"`
	Icon        string    `json:"icon,omitempty" gorm:"type:varchar(100)"`
	Status      string    `json:"status,omitempty" gorm:"type:varchar(100)"`
	UserCreated uint      `json:"userCreated,omitempty" gorm:"not null" binding:"required"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	CreatedAt   time.Time `json:"createdAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
}
