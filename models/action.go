package models

import "time"

type ActionType string

const (
	ActionMoveLeft   ActionType = "MOVE_LEFT"
	ActionMoveRight  ActionType = "MOVE_RIGHT"
	ActionMoveUp     ActionType = "MOVE_UP"
	ActionMoveDown   ActionType = "MOVE_DOWN"
	ActionShootLeft  ActionType = "SHOOT_LEFT"
	ActionShootRight ActionType = "SHOOT_RIGHT"
	ActionShootUp    ActionType = "SHOOT_UP"
	ActionShootDown  ActionType = "SHOOT_DOWN"
	ActionUnknown    ActionType = "UNKNOWN"
)

type ActionDTO struct {
	ID              int64      `json:"id"`
	BotID           int64      `json:"botId"`
	Action          ActionType `json:"action"`
	Forbidden       bool       `json:"forbidden"`
	ExecutionTimeMs int64      `json:"executionTimeMs"`
	CreatedDate     time.Time  `json:"createdDate"`
}
