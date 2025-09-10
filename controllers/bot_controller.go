package controllers

import (
	"net/http"
	"template-go/models"
	"template-go/services"

	"github.com/gin-gonic/gin"
)

type BotController struct {
	botService *services.BotService
}

func NewBotController(botService *services.BotService) *BotController {
	return &BotController{
		botService: botService,
	}
}

func (c *BotController) GetBotInfo(ctx *gin.Context) {
	botInfo, err := c.botService.GetBotInfo()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, botInfo)
		return
	}

	ctx.JSON(http.StatusOK, botInfo)
}

func (c *BotController) Play(ctx *gin.Context) {
	var gameBotDTO models.GameBotDTO

	if err := ctx.ShouldBindJSON(&gameBotDTO); err != nil {
		response := &models.PlayResponseDTO{Action: models.ActionUnknown}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	playResponse, err := c.botService.Play(&gameBotDTO)
	if err != nil {
		response := &models.PlayResponseDTO{Action: models.ActionUnknown}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusOK, playResponse)
}
