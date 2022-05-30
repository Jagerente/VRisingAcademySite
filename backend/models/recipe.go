package models

const (
	RecipeResultField string = "resultItemId"
)

type Recipe struct {
	Entity
	Stations    []string           `json:"stations"`
	Results     []RecipeResult     `json:"results"`
	Time        int32              `json:"time"`
	Ingredients []RecipeIngredient `json:"ingredients"`
	Knowledge   *int               `json:"knowledge"`
}

type RecipeResult struct {
	Id     int32 `json:"itemId"`
	Amount int32 `json:"amount"`
}

type RecipeIngredient struct {
	Id     int32 `json:"itemId"`
	Amount int32 `json:"amount"`
}
