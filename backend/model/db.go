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
		Name:     "根本 洋次郎",
		Email:    "yoni@example.com",
		Password: password,
		IsAdmin:  true,
	}
	DB.FirstOrCreate(&admin, User{Email: admin.Email})

	teamMembers := []User{
		{Name: "生田 まひろ", Email: "ikuta@example.com", Password: password, AdminID: &admin.ID},
		{Name: "根本 洋次郎", Email: "nemoto@example.com", Password: password, AdminID: &admin.ID},
		{Name: "根本 捷太郎", Email: "nemoto@example.com", Password: password, AdminID: &admin.ID},
		{Name: "林 琢磨", Email: "hayasi@example.com", Password: password, AdminID: &admin.ID},
		{Name: "楠本 泰史", Email: "kusumoto@example.com", Password: password, AdminID: &admin.ID},
		{Name: "桑原 将司", Email: "kuwahara@example.com", Password: password, AdminID: &admin.ID},
		{Name: "牧 秀吾", Email: "maki@example.com", Password: password, AdminID: &admin.ID},
		{Name: "宮崎 敏郎", Email: "miyazaki@example.com", Password: password, AdminID: &admin.ID},
		{Name: "山本 雄大", Email: "yamamoto@example.com", Password: password, AdminID: &admin.ID},
		{Name: "知野 直人", Email: "chino@example.com", Password: password, AdminID: &admin.ID},
		{Name: "関根 大気", Email: "sekine@example.com", Password: password, AdminID: &admin.ID},
		{Name: "伊勢 大夢", Email: "ise@example.com", Password: password, AdminID: &admin.ID},
		{Name: "森原 公平", Email: "morihara@example.com", Password: password, AdminID: &admin.ID},
		{Name: "藤田 和也", Email: "hujita@example.com", Password: password, AdminID: &admin.ID},
	}
	for _, member := range teamMembers {
		DB.FirstOrCreate(&member, User{Email: member.Email})
	}
}
