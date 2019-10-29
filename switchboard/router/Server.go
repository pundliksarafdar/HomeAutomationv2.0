package router

import (
	"agent/managers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"gopkg.in/natefinch/lumberjack.v2"
)

func StartWebServer(){
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	if LOG_FILE_LOCATION != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LOG_FILE_LOCATION,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}

	router := gin.Default()
	switchApi := router.Group("/switch")

	switchApi.GET("/ips",(&managers.SwitchManager{}).GetSwitchIps)
	switchApi.GET("/info",(&managers.SwitchManager{}).GetSwitchData)

	//Switch toggle api
	switchApi.POST("/toggle",(&managers.SwitchManager{}).ToggleSwitch)

	//Analytics
	analyticsApi := router.Group("/analytics")
	analyticsApi.GET("/history",(&managers.AnalyticsManager{}).GetHistoricalData)
	analyticsApi.GET("/dashboard",(&managers.AnalyticsManager{}).GetCurrentStatus)
	router.Run(":8080")

}
