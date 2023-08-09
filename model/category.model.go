package model

import "time"

type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ParentId    int       `json:"parentId" gorm:"not null"`
	Order       int       `json:"order" gorm:"not null"`
	Label       string    `json:"label" gorm:"type:varchar(100); not null"`
	Icon        string    `json:"icon,omitempty" gorm:"type:varchar(100)"`
	Status      string    `json:"status,omitempty" gorm:"type:varchar(100)"`
	UserCreated uint      `json:"userCreated,omitempty" gorm:"not null"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	CreatedAt   time.Time `json:"createdAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
}
