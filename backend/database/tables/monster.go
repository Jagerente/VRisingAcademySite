package tables

const (
	MonsterTableName         string = "Monsters"
	MonsterLocationTableName string = "MonsterLocations"
	MonsterDropTableName     string = "MonsterDrops"
	IsBossFieldName          string = "isBoss"
	MonsterIdFieldName       string = "monsterId"
	LocationIdFieldName      string = "locationId"
)

type MonsterDataModel struct {
	GameEntityDataModel
	IsBoss bool
}

type MonsterDropDataModel struct {
	MonsterId int32
	ItemId    int32
}

type MonsterLocationDataModel struct {
	MonsterId  int32
	LocationId int32
}
