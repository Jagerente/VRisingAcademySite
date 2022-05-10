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

type RecipeItemDataModel struct {
	RecipeId int32
	ItemId   int32
}
