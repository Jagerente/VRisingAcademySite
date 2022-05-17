package models

const (
	ItemsTable    string = "Items"
	ItemTierField string = "tier"
)

type Item struct {
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
	Item
	Type       int32    `json:"type"`
	Durability int32    `json:"durability"`
	GearLevel  int32    `json:"gearLevel"`
	MainStat   float64  `json:"mainStat"`
	BonusStats []string `json:"bonusStats"`
	SetId      *int32   `json:"setId"`
}

type Armour struct {
	Item
	Type       int32    `json:"type"`
	Durability int32    `json:"durability"`
	GearLevel  int32    `json:"gearLevel"`
	MainStat   float64  `json:"mainStat"`
	BonusStats []string `json:"bonusStats"`
	SetId      *int32   `json:"setId"`
	SlotId     int32    `json:"slotId"`
}

type Consumable struct {
	Item
	Type int32 `json:"type"`
}

type Reagent struct {
	Item
	Type int32 `json:"type"`
}
