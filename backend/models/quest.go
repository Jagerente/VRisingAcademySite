package models

type Quest struct {
	GameEntity
	Knowledge *int     `json:"knowledge"`
	Goals     []string `json:"goals"`
}
