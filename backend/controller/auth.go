package controller

import (
	"fmt"
	"net/http"
	"time"

	"backend/crypto"
	"backend/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	ID    uint `json:"id"`
	Admin bool `json:"admin"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	studentId := c.FormValue("studentId")
	password := c.FormValue("password")

	fmt.Println(studentId)
	fmt.Println(password)

	user := model.User{}
	result := model.DB.Find(&user, "student_id = ?", studentId)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User not found"})
	}
	if crypto.CompareHashAndPassword(user.Password, password) != nil {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		user.ID,
		user.IsAdmin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func CurrentUser(c echo.Context) model.User {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JwtCustomClaims)
	id := claims.ID
	user := model.User{}
	model.DB.Preload("Team").First(&user, id)
	return user
}
