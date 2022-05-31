package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type ItemJson struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tier        int32  `json:"tier"`
	Type        int32  `json:"type"`
}

type ListResponse struct {
}

func getItemTypesList(ctx *gin.Context) {

	var items = make([]models.ItemType, 0)
	connection := database.CreateConnection()
	defer connection.Close()

	query := `select * from itemtypes`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.ItemType{}

		readError := rows.Scan(
			&item.Id,
			&item.Title)

		if readError != nil {
			fmt.Println(readError)
			continue
		}

		items = append(items, item)
	}
	ctx.JSON(200, items)
}

func getItemsList(ctx *gin.Context) {

	var response = map[string][]models.NewItem{}
	var setMap = map[int32]string{
		-1: "Unset"}

	connection := database.CreateConnection()
	defer connection.Close()

	newQuery := `select
    items.id,
    items.name,
    items.description,
    items.tier,
    items.type as typeid,
    itemtypes.title as typetitle,
    items.knowledgeid,
    array(
        (
            select
                stcns.id
            from
                recipestations as rcst
                join stations as stcns on rcst.stationId = stcns.id
            where
                rcst.recipeId = rcp.id
        )
    ) as stations,
    array(
        (
            select
                recipes.id
            from
                reciperesults
                join recipes on recipes.id = reciperesults.recipeid
            where
                reciperesults.itemid = items.id
        )
    ) as recipes,
    array(
        (
            select
                rc.id
            from
                recipeingredients as rci
                join recipes as rc on rci.recipeId = rc.id
            where
                rci.itemId = items.id
        )
    ) as reagentFor,
    array(
        (
            select
                tgs.value
            from
                itemtags as itt
                join tags as tgs on itt.tagId = tgs.id
            where
                itt.itemId = items.id
        )
    ) as tags,
    stats.durability,
    stats.gearLevel,
    stats.mainStat,
    stats.setid,
    sets.name as setname,
    stats.slotid,
    array(
        (
            select
                secondarystats.bonus
            from
                secondaryitemstats
                join secondarystats on secondarystats.id = secondaryitemstats.secondarystatid
            where
                secondaryitemstats.statsId = stats.id
        )
    ) as bonusStats,
    array(
        (
            select
                lc.name
            from
                itemlocations as itl
                join locations as lc on itl.locationId = lc.id
            where
                itl.itemId = items.id
        )
    ) as locations
from
    items
    join itemtypes on itemtypes.id = items.type
    full join recipeingredients on recipeingredients.itemid = items.id
    full join recipes as rcp on recipeingredients.itemid = rcp.id 
    full join itemstats as stats on stats.id = items.id
    full join sets on sets.id = stats.setid
group by
    items.id,
    itemtypes.title,
    rcp.id,
    stats.id,
    stats.durability,
    stats.gearLevel,
    stats.mainStat,
    stats.setid,
    sets.name
order by
    items.id,
    items.type`

	sets := getSetsList()
	for _, item := range sets {
		if _, ok := setMap[item.Id]; !ok {
			setMap[item.Id] = item.Name
		}
	}

	for _, value := range setMap {
		response[value] = make([]models.NewItem, 0)
	}

	rows, err := connection.Query(newQuery)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		item := models.NewItem{
			Stations:   make([]int32, 0),
			Recipes:    make([]int32, 0),
			ReagentFor: make([]int32, 0),
			Tags:       make([]string, 0),
			BonusStats: make([]string, 0),
			Locations:  make([]string, 0)}

		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.Tier,
			&item.Type,
			&item.TypeName,
			&item.KnowledgeId,
			pq.Array(&item.Stations),
			pq.Array(&item.Recipes),
			pq.Array(&item.ReagentFor),
			pq.Array(&item.Tags),
			&item.Durability,
			&item.GearLevel,
			&item.MainStat,
			&item.SetId,
			&item.Set,
			&item.SlotId,
			pq.Array(&item.BonusStats),
			pq.Array(&item.Locations))

		if readError != nil {
			fmt.Print(readError)
			continue
		}
		var key int32 = -1
		if item.SetId != nil {
			key = *item.SetId
		}
		response[setMap[key]] = append(response[setMap[key]], item)
	}

	ctx.JSON(200, response)
}

func HandleItemsRequest(ctx *gin.RouterGroup) {
	ctx.GET("/list", getItemsList)
	ctx.GET("/types", getItemTypesList)
}
