package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

const (
	MonsterSelectionQuery string = `select
    monsters.id,
    monsters.name,
    monsters.description,
    monsters.factionid,
    factions.name,
    monsters.bloodtypeid,
    bloodtypes.name,
    monsters.level,
    monsters.knowledgeid,
    monsters.mapgenieid,
    array_remove(array_agg(locations.id), NULL)  as locationids,
    array_remove(array_agg(locations.name), NULL)  as locationnames,
    array_remove(array_agg(monsterdrops.itemid), NULL)  as itemids,
    array_remove(array_agg(monsterdrops.droprate), NULL)  as droprates,
    array_remove(array_agg(monsterdrops.amount), NULL)  as amounts
from
    monsters
    join factions on factions.id = monsters.factionid
    join bloodtypes on bloodtypes.id = monsters.bloodtypeid
    left join monsterlocations on monsterlocations.monsterid = monsters.id
    left join locations on monsterlocations.locationid = locations.id
    left join monsterdrops on monsterdrops.monsterid = monsters.id
group by
    monsters.id,
    factions.name,
    bloodtypes.name
order by
    monsters.id`
)

func getFactionList() []models.MonsterFactionObject {
	var items = make([]models.MonsterFactionObject, 0)
	connection := database.CreateConnection()
	defer connection.Close()
	rows, err := connection.Query("select * from factions order by factions.id")

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.MonsterFactionObject{}

		readError := rows.Scan(
			&item.Id,
			&item.Name)

		if readError != nil {
			fmt.Println(readError)
			continue
		}

		items = append(items, item)
	}
	return items
}

func getMonsterList() []models.Monster {
	var items = make([]models.Monster, 0)
	connection := database.CreateConnection()
	defer connection.Close()
	rows, err := connection.Query(MonsterSelectionQuery)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		item := models.Monster{
			Faction:   models.MonsterFactionObject{},
			BloodType: models.MonsterBloodTypeObject{},
			Locations: make([]models.MonsterLocationObject, 0),
			Loot:      make([]models.MonsterDropObject, 0)}

		locationIds := make([]int32, 0)
		locationNames := make([]string, 0)
		dropIds := make([]int32, 0)
		dropRates := make([]float64, 0)
		dropAmounts := make([]int32, 0)

		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.Faction.Id,
			&item.Faction.Name,
			&item.BloodType.Id,
			&item.BloodType.Name,
			&item.Level,
			&item.KnowledgId,
			&item.MapgenieId,
			pq.Array(&locationIds),
			pq.Array(&locationNames),
			pq.Array(&dropIds),
			pq.Array(&dropRates),
			pq.Array(&dropAmounts))

		if readError != nil {
			fmt.Println(readError)
			continue
		}

		for index := range locationIds {
			item.Locations = append(item.Locations, models.MonsterLocationObject{
				Id:   locationIds[index],
				Name: locationNames[index]})
		}

		for index := range dropIds {
			item.Loot = append(item.Loot, models.MonsterDropObject{
				ItemId:   dropIds[index],
				DropRate: dropRates[index],
				Amount:   dropAmounts[index]})
		}

		items = append(items, item)
	}
	return items
}

func handleFactionList(context *gin.Context) {
	items := getFactionList()
	context.JSON(200, items)
}

func handleMonsterList(context *gin.Context) {
	items := getMonsterList()
	context.JSON(200, items)
}

func HandleMonsterRequest(group *gin.RouterGroup) {
	group.GET("/list", handleMonsterList)
	group.GET("factions", handleFactionList)
}
