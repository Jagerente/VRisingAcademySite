package models

const (
	RecipeResultField string = "resultItemId"
)

type Recipe struct {
	GameEntity
	Result      int32              `json:"result"`
	Time        int32              `json:"time"`
	Ingredients []RecipeIngredient `json:"ingredients"`
}

type RecipeIngredient struct {
	Id     int32 `json:"itemId"`
	Amount int32 `json:"amount"`
}
