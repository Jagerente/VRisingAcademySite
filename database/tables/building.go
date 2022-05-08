package tables

const (
	BuildingsTableName  string = "Buildings"
	BuildingNameField   string = "name"
	BuildingDescription string = "description"
)

type BuildingDataModel struct {
	GameEntityDataModel
}
