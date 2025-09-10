package models

type CaseContent string

const (
	CaseContentFirstBot    CaseContent = "FIRST_BOT"
	CaseContentSecondBot   CaseContent = "SECOND_BOT"
	CaseContentBulletLeft  CaseContent = "BULLET_LEFT"
	CaseContentBulletRight CaseContent = "BULLET_RIGHT"
	CaseContentBulletUp    CaseContent = "BULLET_UP"
	CaseContentBulletDown  CaseContent = "BULLET_DOWN"
)

type CaseDTO struct {
	Coordinates *CoordDTO   `json:"coordinates"`
	Content     CaseContent `json:"content"`
	Items       []ItemDTO   `json:"items"`
}
