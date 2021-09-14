package controllers

import (
	"net/http"
	"wm/logger/db"

	"github.com/gin-gonic/gin"
)

func Addlog(c *gin.Context) {
	apiLog := db.ApiLog{
		Api:     c.PostForm("api"),
		Status:  c.PostForm("status"),
		Latency: c.PostForm("latency"),
		Method:  c.PostForm("method"),
	}
	db.DB.Create(&apiLog)
	var apiCount db.ApiCount

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
