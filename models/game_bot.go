package models

type GameBotDTO struct {
	Game *GameDTO `json:"game"`
	Bot  *BotDTO  `json:"bot"`
}
