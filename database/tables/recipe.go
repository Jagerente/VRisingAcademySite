package tables

const (
	RecipeTableName     string = "Recipes"
	RecipeItemTableName string = "RecipeIngredients"
	RecipeResultField   string = "resultItemId"
	RecipeIdFieldName   string = "recipeId"
)

type RecipeDataModel struct {
	GameEntityDataModel
	ResultId int32
}

type RecipeIngridientDataModel struct {
	RecipeId int32
	ItemId   int32
	Amount   int32
}
