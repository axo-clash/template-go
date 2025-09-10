package models

type BotState string

const (
	BotStateUnknown            BotState = "UNKNOWN"
	BotStateReady              BotState = "READY"
	BotStatePlaying            BotState = "PLAYING"
	BotStateWaitingForOpponent BotState = "WAITING_FOR_OPPONENT"
)

type BotDataDTO struct {
	Life             int32    `json:"life"`
	ForbiddenActions int32    `json:"forbiddenActions"`
	State            BotState `json:"state"`
	LastActionTimeMs int64    `json:"lastActionTimeMs"`
}
