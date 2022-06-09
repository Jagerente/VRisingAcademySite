package models

type BloodType struct {
	GameEntity
	Bonuses []string `json:"bonuses"`
}
