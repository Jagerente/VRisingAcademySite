package models

type Structure struct {
	GameEntity
	Variants  []int32  `json:"variants"`
	Type      int32    `json:"type"`
	Tier      int      `json:"tier"`
	Tags      []string `json:"tags"`
	Recipes   []int32  `json:"recipes"`
	IsVariant bool     `json:"isVariant"`
}
