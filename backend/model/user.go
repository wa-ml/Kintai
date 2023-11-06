package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"not null"`
	StudentId string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	IsAdmin   bool
}
