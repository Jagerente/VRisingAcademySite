package models

type SalvageableResult struct {
	ItemId int32 `json:"itemId"`
	Amount int32 `json:"amount"`
}

type Salvageable struct {
	Entity
	ItemId  int32               `json:"itemId"`
	Results []SalvageableResult `json:"results"`
}
