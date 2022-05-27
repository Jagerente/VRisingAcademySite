package models

type Knowledge struct {
	Entity
	Recipes []int `json:"recipes"`
	Spells  []int `json:"spells"`
	Quest   *int  `json:"quest"`
	Monster *int  `json:"boss"`
}
