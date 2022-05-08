package tables

const (
	TagsTableName string = "ItemTags"
	TagValueField string = "value"
)

type TagDataModel struct {
	EntityDataModel
	Value string
}
