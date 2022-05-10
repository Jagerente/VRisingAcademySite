package tables

const (
	ItemTableName    string = "Items"
	ItemTierField    string = "tier"
	ItemIdFieldName  string = "itemId"
	TagIdFieldName   string = "tagId"
	ItemTagTableName string = "ItemTags"
)

type ItemDataModel struct {
	GameEntityDataModel
	Tier int32
}

type ItemTagDataModel struct {
	ItemId int32
	TagId  int32
}
