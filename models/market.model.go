package models

import (
	"time"
)

type Market struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"market"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
