package main

import (
	"backend/controller"
	"backend/model"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	db, _ := model.DB.DB()
	defer db.Close()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/user", controller.CreateAdminUser)
	e.POST("/login", controller.Login)
	/// Restricted group
	r := e.Group("/auth")

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(controller.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("/users", controller.GetUsers)
	r.POST("/user", controller.CreateUser)
	r.GET("/check", controller.CheckAdmin)
	r.GET("/user", controller.GetUser)
	r.PUT("/user", controller.UpdateUser)
	// r.DELETE("/users/:id", controller.DeleteUser)

	r.POST("/kintaiLog", controller.CreateKintaiLog)
	r.GET("/kintaiLog", controller.GetKintaiLogs)

	e.Logger.Fatal(e.Start(":8080"))
}
