package main

import (
	"backend/controller"
	"backend/model"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, _ := model.DB.DB()
	defer db.Close()
	e.POST("/login", controller.Login)
	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
	e.Logger.Fatal(e.Start(":8080"))
}
