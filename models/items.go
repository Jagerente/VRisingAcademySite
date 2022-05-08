package models

const (
	ItemsTable    string = "Items"
	ItemTierField string = "tier"
)

type Item struct {
	Entity
	Tags []Tag `json:"tags"`
	Tier int32 `json:"tier"`
}
