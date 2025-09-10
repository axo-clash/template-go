package services

import (
	"math/rand"
	"template-go/models"
	"time"
)

type BotService struct{}

func NewBotService() *BotService {
	return &BotService{}
}

func (s *BotService) GetBotInfo() (*models.BotInfoDTO, error) {
	return &models.BotInfoDTO{
		Name:    "Go Template Bot",
		Version: "1.0.0",
	}, nil
}

func (s *BotService) Play(gameBotDTO *models.GameBotDTO) (*models.PlayResponseDTO, error) {
	// List of all possible actions
	actions := []models.ActionType{
		models.ActionMoveLeft,
		models.ActionMoveRight,
		models.ActionMoveUp,
		models.ActionMoveDown,
		models.ActionShootLeft,
		models.ActionShootRight,
		models.ActionShootUp,
		models.ActionShootDown,
	}

	// Initialize random seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Select random action
	randomAction := actions[r.Intn(len(actions))]

	return &models.PlayResponseDTO{
		Action: randomAction,
	}, nil
}
