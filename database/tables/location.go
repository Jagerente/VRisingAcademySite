package tables

const (
	LocationTableName string = "Locations"
)

type LocationDataModel struct {
	EntityDataModel
	Value string
}
