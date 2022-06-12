package models

import "database/sql"

type MonsterFactionObject struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type MonsterBloodTypeObject struct {
	Entity
	Name string `json:"name"`
}

type MonsterLocationObject struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type MonsterDropObject struct {
	ItemId   int32   `json:"itemId"`
	DropRate float64 `json:"dropRate"`
	Amount   int32   `json:"amount"`
}

type Monster struct {
	GameEntity
	Faction    MonsterFactionObject    `json:"faction"`
	BloodType  MonsterBloodTypeObject  `json:"bloodType"`
	Level      int32                   `json:"level"`
	KnowledgId sql.NullInt32           `json:"knowledgeId"`
	MapgenieId sql.NullInt32           `json:"mapgenieId"`
	Locations  []MonsterLocationObject `json:"locations"`
	Loot       []MonsterDropObject     `json:"loot"`
}
