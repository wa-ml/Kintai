package model

import (
	"backend/crypto"
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
	DB.AutoMigrate(&KintaiLog{})

	//初期データ投入
	password, _ := crypto.PasswordEncrypt("hogehoge")

	admin := User{
		Name:     "Admin User",
		Email:    "admin@example.com",
		Password: password,
		IsAdmin:  true,
	}
	DB.FirstOrCreate(&admin, User{Email: admin.Email})

	teamMembers := []User{
		{Name: "User One", Email: "user1@example.com", Password: password, AdminID: &admin.ID},
		{Name: "User Two", Email: "user2@example.com", Password: password, AdminID: &admin.ID},
	}
	for _, member := range teamMembers {
		DB.FirstOrCreate(&member, User{Email: member.Email})
	}
}
