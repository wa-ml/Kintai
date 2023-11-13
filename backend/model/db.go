package model

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "user:password@tcp(db_container:3306)/kintai-db?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"

	const maxAttempts = 10
	for i := 0; i < maxAttempts; i++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Attempt %d: database can't connect, error: %v", i+1, err)
		time.Sleep(time.Second * time.Duration(2<<i))
	}

	if err != nil {
		log.Fatalf("After %d attempts, last error: %v", maxAttempts, err)
	}

	DB.AutoMigrate(&User{})
}
