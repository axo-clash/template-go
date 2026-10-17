package models

type TrapDTO struct {
	PlacedBy         BotType `json:"placedBy"`
	PlacedAtTurn     int     `json:"placedAtTurn"`
	ArmedFromTurn    int     `json:"armedFromTurn"`
	ExpiresAfterTurn int     `json:"expiresAfterTurn"`
}
