package tables

const (
	ItemsTableName string = "Items"
	ItemTierField  string = "tier"
)

type ItemDataModel struct {
	GameEntityDataModel
	Tier int32
}
