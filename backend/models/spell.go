package models

type SpellTypeObject struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type SpellSchoolObject struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type Spell struct {
	GameEntity
	SchoolId   int32             `json:"-"`
	SchoolName string            `json:"-"`
	School     SpellSchoolObject `json:"school"`
	TypeId     int32             `json:"-"`
	TypeName   string            `json:"-"`
	Type       SpellTypeObject   `json:"type"`
	Cooldown   float64           `json:"cooldown"`
	CastTime   float64           `json:"castTime"`
	Charges    int32             `json:"charges"`
	Knowledge  *int32            `json:"knowledge"`
}

type SpellWithSchoolId struct {
	Spell
	SchoolId int32 `json:"schoolId"`
}

type SpellType struct {
	Entity
	Name string `json:"title"`
}

type SpellSchool struct {
	GameEntity
}
