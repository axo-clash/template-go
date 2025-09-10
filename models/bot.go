package models

type BotType string

const (
	BotTypeFirst  BotType = "FIRST_BOT"
	BotTypeSecond BotType = "SECOND_BOT"
)

type BotDTO struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Version     string      `json:"version"`
	URL         string      `json:"url"`
	Data        *BotDataDTO `json:"data"`
	BotType     BotType     `json:"botType"`
	Coordinates *CoordDTO   `json:"coordinates"`
	Faster      bool        `json:"faster"`
}
