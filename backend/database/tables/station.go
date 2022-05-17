package tables

const (
	StationsTableName = "stations"
)

type StationDataModel struct {
	GameEntityDataModel
	LocationId int32
}

type StationRecipesDataModel struct {
	StationId int32
	RecipeId  int32
}
