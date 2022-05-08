package models

const (
	RecipeResultField string = "resultItemId"
)

type Recipe struct {
	Entity
	Result      Item   `json:"result"`
	Ingredients []Item `json:"ingredients"`
}
