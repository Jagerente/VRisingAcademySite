package tables

const (
	TagsTableName string = "Tags"
	TagValueField string = "value"
)

type TagDataModel struct {
	EntityDataModel
	Value string
}
