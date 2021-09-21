package controllers

import (
	"net/http"
	"wm/logger/db"

	"github.com/gin-gonic/gin"
)

// 로그 저장 및 서비스 별 호출 횟수 저장
func Addlog(c *gin.Context) {
	apiLog := db.ApiLog{
		Api:     c.PostForm("api"),
		Status:  c.PostForm("status"),
		Latency: c.PostForm("latency"),
		Method:  c.PostForm("method"),
	}
	// Gorm을 사용한 데이터 저장
	db.DB.Create(&apiLog)
	var apiCount db.ApiCount

	// Gorm을 사용하여 서비스 별 API 호출 횟수 저장
	if err := db.DB.Where("api = ? AND method = ?", apiLog.Api, apiLog.Method).First(&apiCount).Error; err != nil {
		db.DB.Create(&db.ApiCount{Api: apiLog.Api, Count: 1, Method: apiLog.Method})
	} else {
		apiCount.Count += 1
		db.DB.Save(&apiCount)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "log saved",
		"body":    apiLog.Api,
	})
}

func ListLogs(c *gin.Context) {
	var apiLogs []db.ApiLog
	if service, check := c.GetQuery("service"); check {
		db.DB.Where("api LIKE ?", "/"+service+"%").Find(&apiLogs)
	} else {
		db.DB.Find(&apiLogs)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "list all Log",
		"body":    apiLogs,
	})
}

func ListCounts(c *gin.Context) {
	var apiCounts []db.ApiCount
	if service, check := c.GetQuery("service"); check {
		db.DB.Where("api LIKE ?", "/"+service+"%").Find(&apiCounts)
	} else {
		db.DB.Find(&apiCounts)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "list all count",
		"body":    apiCounts,
	})
}

func DeleteWrongLog(c *gin.Context) {
	db.DB.Unscoped().Where("Method IS NULL").Delete(&db.ApiLog{})
	db.DB.Unscoped().Where("Method = ", " ").Delete(&db.ApiLog{})
	db.DB.Unscoped().Where("Latency LIKE ?", "%ms").Delete(&db.ApiLog{})
	db.DB.Unscoped().Where("Latency LIKE ?", "%µs").Delete(&db.ApiLog{})

	db.DB.Unscoped().Where("Method IS NULL").Delete(&db.ApiCount{})
	db.DB.Unscoped().Where("Method = ", " ").Delete(&db.ApiCount{})
	db.DB.Unscoped().Where("Api = ", " ").Delete(&db.ApiCount{})

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}
