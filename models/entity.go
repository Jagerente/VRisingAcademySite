package models

const (
	IdField string = "id"
)

type Entity struct {
	Id int32 `json:"id"`
}

func GetId(en *Entity) int32 {
	return en.Id
}
