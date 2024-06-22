package core

import (
	"time"
)

type BaseModel struct {
	ID        uint       `json:"id" example:"1" gorm:"primarykey" format:"int64" description:"User ID" `
	CreatedAt time.Time  `json:"createdAt" example:"2022-01-01T00:00:00Z" format:"date-time" description:"Creation timestamp"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" valid:"-" `
} //@name models.BaseModel
