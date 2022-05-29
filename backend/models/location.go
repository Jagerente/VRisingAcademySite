package models

type Location struct {
	GameEntity
	Region     string  `json:"region"`
	Loot       []int32 `json:"loot"`
	MapgenieId int32   `json:"mapgenieId"`
}

type LocationWithRegion struct {
	Location
	RegionId int32 `json:"regionId"`
}

type Region struct {
	Entity
	Name string `json:"name"`
}