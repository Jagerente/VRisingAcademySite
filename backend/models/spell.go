package models

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
