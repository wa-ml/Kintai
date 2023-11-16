package controller

import (
	"net/http"
	"backend/model"
	"github.com/labstack/echo/v4"
	"time"
)

type LogTime struct {
	ArrivalTime time.Time `json:"arrival_time"`
	ReturnTime 	time.Time `json:"return_time"`
}

type ResJson struct {
	LogTimes	[]LogTime
	TotalTime	float64
}

// for Admin
func CreateKintaiLog(c echo.Context) error {
	user := CurrentUser(c)

	kintaiLog := model.KintaiLog{}
	if err := c.Bind(&kintaiLog); err != nil {
		return err
	}

	kintaiLog.UserID = user.ID

	kintaiLogs := []model.KintaiLog{}
	
	if err := model.DB.Find(&kintaiLogs, "user_id = ?", user.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	nowStatus := ""
	for _, log := range kintaiLogs {
		nowStatus = log.Status
	}

	if nowStatus == "Active" {
		kintaiLog.Status = "NonActive"
	} else {
		kintaiLog.Status = "Active"
	}

	jst := time.FixedZone("Asia/Tokyo", 9*60*60);
	nowJST := time.Now().UTC().In(jst)
	kintaiLog.LogTime = nowJST

	if err := model.DB.Create(&kintaiLog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, kintaiLog)
}

func GetKintaiLogs(c echo.Context) error {
	user := CurrentUser(c)

	var res ResJson


	var logs []LogTime
	kintaiLogs := []model.KintaiLog{}
	
	if err := model.DB.Find(&kintaiLogs, "user_id = ?", user.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var logTime LogTime
	var dif time.Duration = 0
	var nanosecondsNow time.Duration
	for _, log := range kintaiLogs {
		if (log.Status == "Active") {
			logTime.ArrivalTime = log.LogTime
		} else {
			logTime.ReturnTime = log.LogTime
			logs = append(logs, logTime)
			dif = logTime.ReturnTime.Sub(logTime.ArrivalTime)
			nanosecondsNow += dif
		}
	}

	nanosecondsPerHour := int64(3600 * 1e9)
	hours := float64(nanosecondsNow) / float64(nanosecondsPerHour)
	res.TotalTime = hours
	res.LogTimes = logs
	
	return c.JSON(http.StatusOK, res)
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
