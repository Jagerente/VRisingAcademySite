package tables

const (
	RecipeTableName   string = "Recipes"
	RecipeResultField string = "resultItemId"
)

type RecipeDataModel struct {
	GameEntityDataModel
	ResultId int32
}
