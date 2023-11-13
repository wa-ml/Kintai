package controller

import (
	"net/http"

	"backend/crypto"
	"backend/model"

	"github.com/labstack/echo/v4"
)

// for Admin
func CreateUser(c echo.Context) error {
	admin_user := CurrentUser(c)
	if !admin_user.IsAdmin {
		return c.String(http.StatusBadRequest, "権限がありません")
	}

	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	user.AdminID = &admin_user.ID
	user.IsAdmin = false

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
	if !user.IsAdmin {
		return c.String(http.StatusBadRequest, "権限がありません")
	}
	return c.JSON(http.StatusOK, user.Team)
}

// for all
func CheckAdmin(c echo.Context) error {
	user := CurrentUser(c)
	return c.JSON(http.StatusOK, user.IsAdmin)
}

func GetUser(c echo.Context) error {
	user := CurrentUser(c)
	return c.JSON(http.StatusOK, user)
}

// func UpdateUser(c echo.Context) error {
// 	user := model.User{}
// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}
// 	model.DB.Save(&user)
// 	return c.JSON(http.StatusOK, user)
// }

// func DeleteUser(c echo.Context) error {
// 	user := model.User{}
// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}
// 	model.DB.Delete(&user)
// 	return c.JSON(http.StatusOK, user)
// }
