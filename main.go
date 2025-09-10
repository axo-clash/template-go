package main

import (
	"template-go/controllers"
	"template-go/services"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	botService := services.NewBotService()
	botController := controllers.NewBotController(botService)

	r.GET("/info", botController.GetBotInfo)
	r.POST("/play", botController.Play)

	r.Run(":6000")
}
