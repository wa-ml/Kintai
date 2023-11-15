package controller

import (
	"net/http"

	"backend/crypto"
	"backend/mail"
	"backend/model"

	"github.com/labstack/echo/v4"
)

type PasswordUpdateRequest struct {
	Password string
}

// For Admin
func CreateAdminUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	user.IsAdmin = true
	user.AdminID = nil
	initPassword := user.Password

	hashedPassword, encryptErr := crypto.PasswordEncrypt(initPassword)
	if encryptErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to hash the password")
	}
	user.Password = hashedPassword

	if err := model.DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

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
	initPassword := user.Password

	hashedPassword, encryptErr := crypto.PasswordEncrypt(initPassword)
	if encryptErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to hash the password")
	}
	user.Password = hashedPassword

	if err := model.DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	mail.SendEmail(user.Email, initPassword)

	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
	user := CurrentUser(c)
	if !user.IsAdmin {
		return c.String(http.StatusBadRequest, "権限がありません")
	}
	return c.JSON(http.StatusOK, user.Team)
}

// For ALL
func CheckAdmin(c echo.Context) error {
	user := CurrentUser(c)
	return c.JSON(http.StatusOK, user.IsAdmin)
}

func GetUser(c echo.Context) error {
	user := CurrentUser(c)
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	user := CurrentUser(c)

	var passwordData PasswordUpdateRequest
	if err := c.Bind(&passwordData); err != nil {
		return err
	}

	hashedPassword, encryptErr := crypto.PasswordEncrypt(passwordData.Password)
	if encryptErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to hash the password")
	}

	result := model.DB.Model(&model.User{}).Where("id = ?", user.ID).Update("password", hashedPassword)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.NoContent(http.StatusOK)
}

// func DeleteUser(c echo.Context) error {
// 	user := model.User{}
// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}
// 	model.DB.Delete(&user)
// 	return c.JSON(http.StatusOK, user)
// }
