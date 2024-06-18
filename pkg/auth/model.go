package auth

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type RefreshTokenStatus string

const (
	New  RefreshTokenStatus = "new"
	Used RefreshTokenStatus = "used"
)

func (st *RefreshTokenStatus) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		*st = RefreshTokenStatus(b)
	}
	return nil
}

func (st RefreshTokenStatus) Value() (driver.Value, error) {
	return string(st), nil
}

type RefreshToken struct {
	gorm.Model
	Jti    string             `gorm:"not null" json:"jti"`
	Parent string             `gorm:"foreignkey:RefreshToken.ID;default:null" json:"parent"`
	Status RefreshTokenStatus `gorm:"not null type:num('new','used');default:'new'" json:"status"`
}
