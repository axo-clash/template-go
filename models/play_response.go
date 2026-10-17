package models

type PlayResponseDTO struct {
	Coordinates *CoordDTO  `json:"coordinates,omitempty"`
	Action      ActionType `json:"action"`
}
