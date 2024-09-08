package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model    `json:"-"`
	Email         string               `gorm:"unique;not null" json:"email"`
	Password      string               `gorm:"not null" json:"-"`
	FullName      string               `gorm:"size:255;default:''" json:"fullname"`
	RefreshTokens []RefreshTokenFamily `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	ID            uint                 `gorm:"primarykey" json:"id"`
}
