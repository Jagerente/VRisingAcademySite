package models

const (
	IdField string = "id"
)

type Entity struct {
	Id int32 `json:"id"`
}

type GameEntity struct {
	Entity
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func GetId(en *Entity) int32 {
	return en.Id
}
