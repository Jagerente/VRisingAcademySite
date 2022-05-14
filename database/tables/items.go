package tables

const (
	ItemTableName    string = "Items"
	ItemTierField    string = "tier"
	ItemIdFieldName  string = "itemId"
	TagIdFieldName   string = "tagId"
	TypeField        string = "type"
	ItemTagTableName string = "ItemTags"
	StationIdField   string = "stationId"
)

type ItemDataModel struct {
	GameEntityDataModel
	Tier int32
	Type int32
}

type ItemTagDataModel struct {
	ItemId int32
	TagId  int32
}

type ItemLocationsDataModel struct {
	ItemId     int32
	LocationId int32
}
