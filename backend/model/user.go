package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	IsAdmin  bool   `gorm:"default:false"`
	AdminID  *uint
	Team     []User `gorm:"foreignkey:AdminID"`
	Logs	 []KintaiLog
}
