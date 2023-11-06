package main

import (
	"net/http"

	"log"

	"github.com/labstack/echo/v4"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type User struct {
	gorm.Model
	Name      string `gorm:"not null"`
	StudentId string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	IsAdmin   bool
}

func init() {
	dsn := "user:password@tcp(db_container:3306)/kintai-db?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + "database can't connect")
	}
	DB.AutoMigrate(&User{})
}

func CreateUser(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func main() {
	e := echo.New()
	db, _ := DB.DB()
	defer db.Close()

	e.POST("/users", CreateUser)
	e.Logger.Fatal(e.Start(":8080"))
}
