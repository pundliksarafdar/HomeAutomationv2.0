package router

import (
	"agent/managers"
	"github.com/gin-gonic/gin"
)

func StartWebServer(){
	router := gin.Default()
	switchApi := router.Group("/switch")

	switchApi.GET("/ips",(&managers.SwitchManager{}).GetSwitchIps)
	switchApi.GET("/info",(&managers.SwitchManager{}).GetSwitchData)

	//Switch toggle api
	switchApi.POST("/toggle",(&managers.SwitchManager{}).ToggleSwitch)
	router.Run(":8080")

}
