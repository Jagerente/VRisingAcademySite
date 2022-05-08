package models

const (
	TagValueField string = "value"
)

type Tag struct {
	Entity
	Value string `json:"value"`
}
