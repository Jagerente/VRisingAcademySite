package models

type FullSpell struct {
	Spell
	SchoolId int32 `json:"schoolId"`
	TypeId   int32 `json:"typeId"`
}

type Spell struct {
	GameEntity
	School    string  `json:"school"`
	Type      string  `json:"type"`
	Cooldown  float64 `json:"cooldown"`
	CastTime  float64 `json:"castTime"`
	Charges   int32   `json:"charges"`
	Knowledge *int32  `json:"knowledge"`
}

type SpellWithSchoolId struct {
	Spell
	SchoolId int32 `json:"schoolId"`
}

type SpellType struct {
	Entity
	Title string `json:"title"`
}

type SpellSchool struct {
	GameEntity
}
