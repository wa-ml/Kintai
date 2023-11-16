package model

import (
	"time"

	"gorm.io/gorm"
)

type KintaiLog struct {
	gorm.Model
	UserID  uint      `gorm:"not null"`
	Status  string    `gorm:"not null"`
	LogTime time.Time `gorm:"autoCreateTime"`
}
