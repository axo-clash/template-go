package models

import (
	"encoding/json"
	"time"
)

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

type UnixTime time.Time

func (t *UnixTime) UnmarshalJSON(data []byte) error {
	var timestamp float64
	if err := json.Unmarshal(data, &timestamp); err != nil {
		return err
	}
	seconds := int64(timestamp)
	nanoseconds := int64((timestamp - float64(seconds)) * 1e9)
	*t = UnixTime(time.Unix(seconds, nanoseconds))
	return nil
}

func (t UnixTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).UnixMilli())
}

type ActionDTO struct {
	ID              int64      `json:"id"`
	BotID           int64      `json:"botId"`
	Action          ActionType `json:"action"`
	Forbidden       bool       `json:"forbidden"`
	ExecutionTimeMs int64      `json:"executionTimeMs"`
	CreatedDate     UnixTime   `json:"createdDate"`
}
