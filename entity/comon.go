package entity

import (
	"time"

	"gorm.io/gorm"
)

type Timestamp struct {
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

type Authorization struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}
