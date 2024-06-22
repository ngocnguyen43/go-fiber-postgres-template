package user

import (
	"go-fiber-postgres-template/pkg/core"
)

// User struct
// swagger:model
type User struct {
	core.BaseModel
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Names    string `json:"names"`
}
