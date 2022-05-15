package models

const (
	RecipeResultField string = "resultItemId"
)

type Recipe struct {
	Entity
	Results     []RecipeResult     `json:"results"`
	Time        int32              `json:"time"`
	Ingredients []RecipeIngredient `json:"ingredients"`
}

type RecipeResult struct {
	Id     int32 `json:"itemId"`
	Amount int32 `json:"amount"`
}

type RecipeIngredient struct {
	Id     int32 `json:"itemId"`
	Amount int32 `json:"amount"`
}
