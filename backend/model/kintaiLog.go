package model

import (
	"gorm.io/gorm"
	"time"
)

type KintaiLog struct {
	gorm.Model
	UserID  uint      `gorm:"not null"`
	Status  string    `gorm:"not null"`
	LogTime time.Time `gorm:"autoCreateTime"`
}
