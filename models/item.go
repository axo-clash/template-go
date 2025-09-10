package models

type ItemType string

const (
	ItemBulletLeft  ItemType = "BULLET_LEFT"
	ItemBulletRight ItemType = "BULLET_RIGHT"
	ItemBulletUp    ItemType = "BULLET_UP"
	ItemBulletDown  ItemType = "BULLET_DOWN"
)

type ItemDTO struct {
	Type  ItemType `json:"type"`
	Owner *BotDTO  `json:"owner"`
}
