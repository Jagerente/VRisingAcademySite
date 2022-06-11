package models

import (
	"database/sql"
)

const (
	ItemsTable    string = "Items"
	ItemTierField string = "tier"
)

type ItemTypeObject struct {
	Entity
	Name string `json:"name"`
}

type ItemSetObject struct {
	Id          int32   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Item struct {
	GameEntity
	Tier          int32            `json:"tier"`
	Type          ItemTypeObject   `json:"type"`
	KnowledgeId   *int32           `json:"knowledge"`
	Stations      []sql.NullInt32  `json:"stations"`
	Recipes       []sql.NullInt32  `json:"recipes"`
	ReagentFor    []sql.NullInt32  `json:"reagentFor"`
	Tags          []string         `json:"tags"`
	Durability    *int32           `json:"durability"`
	GearLevel     *int32           `json:"gearLevel"`
	MainStat      *float64         `json:"mainStat"`
	BonusStats    []sql.NullString `json:"bonusStats"`
	SetId         *int32           `json:"-"`
	Set           *ItemSetObject   `json:"set"`
	SlotId        *int32           `json:"slot"`
	Locations     []sql.NullInt32  `json:"locations"`
	Variants      []sql.NullInt32  `json:"variants"`
	Salvageables  []sql.NullInt32  `json:"salvageables"`
	SalvageableOf []sql.NullInt32  `json:"salvageableOf"`
}

type OldItem struct {
	GameEntity
	Tier       int32    `json:"tier"`
	Tags       []string `json:"tags"`
	Recipes    []int32  `json:"recipes"`
	ReagentFor []int32  `json:"reagentFor"`
	Stations   []int32  `json:"stations"`
	//Salvageable []SalvageableItem `json:"salavageable"`
	Locations []string `json:"locations"`
}

type SalvageableItem struct {
	Id     int32 `json:"id"`
	Amount int32 `json:"amount"`
}

type Weapon struct {
	OldItem
	Type       int32    `json:"type"`
	Durability int32    `json:"durability"`
	GearLevel  int32    `json:"gearLevel"`
	MainStat   float64  `json:"mainStat"`
	BonusStats []string `json:"bonusStats"`
	SetId      *int32   `json:"setId"`
}

type Armour struct {
	OldItem
	Type       int32    `json:"type"`
	Durability int32    `json:"durability"`
	GearLevel  int32    `json:"gearLevel"`
	MainStat   float64  `json:"mainStat"`
	BonusStats []string `json:"bonusStats"`
	SetId      *int32   `json:"setId"`
	SlotId     int32    `json:"slotId"`
}

type Consumable struct {
	OldItem
	Type int32 `json:"type"`
}

type Reagent struct {
	OldItem
	Type int32 `json:"type"`
}
