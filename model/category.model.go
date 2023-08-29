package model

import (
	"time"

	"github.com/stoewer/go-strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Label       string    `json:"label" gorm:"type:varchar(100);unique;not null" binding:"required"`
	Order       int       `json:"order" gorm:"not null"`
	Param       string    `json:"param" gorm:"type:varchar(100);not null"`
	Icon        string    `json:"icon,omitempty" gorm:"type:varchar(100)"`
	Status      string    `json:"status,omitempty" gorm:"type:varchar(100)"`
	UserCreated uint      `json:"userCreated,omitempty" gorm:"not null" binding:"required"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	CreatedAt   time.Time `json:"createdAt" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
}

func (o *Category) SetRouterLink() {
	o.Label = cases.Title(language.Und).String(o.Label)
	o.Param = strcase.KebabCase(o.Label)
}
