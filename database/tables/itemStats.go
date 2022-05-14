package tables

const (
	ItemStatsTableName string = "ItemStats"
	MainStatField      string = "mainStat"
	SetField           string = "set"
	SetBonusField      string = "setBonus"
	BonusField         string = "bonus"
	StatsIdField       string = "statsId"
	SecondaryStat      string = "seconderyStat"
	GearLevelField     string = "gearLevel"
	DurabilityField    string = "durability"
)

type ItemStatsDataModel struct {
	EntityDataModel
	MainStat   float64
	Set        string
	SetBonus   string
	GearLevel  int32
	Durability int32
}

type SecondaryStatsDataModel struct {
	EntityDataModel
	Bonus string
}

type ItemSecondaryStatsDataModel struct {
	StatsId          int32
	SecondaryStatsId int32
}
