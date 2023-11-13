package controller

import (
	"fmt"
	"net/http"

	"backend/crypto"
	"backend/model"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	hashedPassword, encryptErr := crypto.PasswordEncrypt(user.Password)
	if encryptErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to hash the password")
	}
	user.Password = hashedPassword

	if err := model.DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
	user := CurrentUser(c)
	return c.JSON(http.StatusOK, user)
}

func GetUser(c echo.Context) error {
	user := model.User{}
	fmt.Println(c.Param("id"))
	result := model.DB.First(&user, c.Param("id"))
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	model.DB.Save(&user)
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	model.DB.Delete(&user)
	return c.JSON(http.StatusOK, user)
}
