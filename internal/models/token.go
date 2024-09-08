package models

import "gorm.io/gorm"

type RefreshTokenStatus string

type RefreshTokenFamilyStatus string

const (
	InUse    RefreshTokenStatus       = "inuse"
	Used     RefreshTokenStatus       = "used"
	Active   RefreshTokenFamilyStatus = "active"
	Inactive RefreshTokenFamilyStatus = "inactive"
)

type RefreshToken struct {
	gorm.Model
	JTI                  string             `gorm:"unique;not null" json:"jti"`
	Status               RefreshTokenStatus `gorm:"default:'inuse'"`
	Parent               *RefreshToken      `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE;"`
	ParentID             uint               `gorm:"default:null"`
	RefreshTokenFamilyID uint               `gorm:"not null"`
	RefreshTokenFamily   RefreshTokenFamily
}

type RefreshTokenFamily struct {
	gorm.Model
	Status        RefreshTokenFamilyStatus `gorm:"default:'active'"`
	UserID        uint                     `gorm:"not null"`
	RefreshTokens []RefreshToken           `gorm:"constraint:OnDelete:CASCADE;"`
}
