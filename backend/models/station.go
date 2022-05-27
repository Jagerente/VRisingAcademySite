package models

type Station struct {
	GameEntity
	Recipes        []int32 `json:"recipes"`
	StationRecipes []int32 `json:"stationRecipes"`
	Knowledges     []int32 `json:"knowledges"`
}
