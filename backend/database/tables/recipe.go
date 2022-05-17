package tables

const (
	RecipeTableName     string = "Recipes"
	RecipeItemTableName string = "RecipeIngredients"
	RecipeResultField   string = "resultItemId"
	RecipeIdFieldName   string = "recipeId"
)

type RecipeDataModel struct {
	Id   int32
	Time int32
}

type RecipeResultDataModel struct {
	RecipeId int32
	ItemId   int32
	Amount   int32
}

type RecipeIngridientDataModel struct {
	RecipeId int32
	ItemId   int32
	Amount   int32
}
