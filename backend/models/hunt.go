package models

type Hunt struct {
	Entity
	Region      string            `json:"region"`
	Location    string            `json:"location"`
	MaxServants int32             `json:"maxServants"`
	Difficulty  int32             `json:"difficulty"`
	BonusId     int32             `json:"bonusBlood"`
	Rewards     []HuntRewardValue `json:"rewards"`
}

type HuntRewardValue struct {
	ItemId        int32  `json:"itemId"`
	Time          int32  `json:"time"`
	Servants      int32  `json:"servantsAmount"`
	Bonus         string `json:"bonus"`
	MinimalAmount int32  `json:"minimalAmount"`
	MaximumAmount int32  `json:"maximumAmount"`
}
