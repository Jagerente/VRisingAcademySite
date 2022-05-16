package tables

type SpellDataModel struct {
	GameEntityDataModel
	SchoolId int32
	TypeId   int32
	Cooldown float64
	CastTime float64
	Charges  int32
}

type SpellSchoolDataModel struct {
	Id   int32
	Name string
}

type SpellTypeDataModel struct {
	Id    int32
	Title string
}
