package model

import (
	"gorm.io/gorm"
)

type Plant struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null;unique"`
	Category    string  `json:"category" gorm:"not null;unique"`
	Price       float64 `json:"price" gorm:"not null;"`
	Description string  `json:"description" gorm:"not null;"`
	Avatar      string  `json:"avatar" gorm:"not null;"`
	UserID      uint    `gorm:"not null;"`
	User        User    `gorm:"not null; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
