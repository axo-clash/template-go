package main

import (
	"log"
	"os"
	"template-go/controllers"
	"template-go/services"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	botService := services.NewBotService()
	botController := controllers.NewBotController(botService)

	r.GET("/api/v1/bots/info", botController.GetBotInfo)
	r.POST("/api/v1/bots/play", botController.Play)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	log.Default().Printf("Server is running on port %s", port)

	r.Run(":" + port)
}
