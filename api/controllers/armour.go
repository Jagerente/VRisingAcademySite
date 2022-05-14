package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func getArmourList(ctx *gin.Context) {

	var items = make([]models.Armour, 0)
	connection := database.CreateConnection()
	defer connection.Close()

	query := `select
    items.id,
    items.name,
    items.description,
    items.tier,
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
    rcp.id as recipeId,
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
    stats.set,
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
                lc.value
            from
                itemlocations as itl
                join locations as lc on itl.locationId = lc.id
            where
                itl.itemId = items.id
        )
    ) as locations
from
    items
    full join recipes as rcp on rcp.resultItemId = items.id
    full join itemstats as stats on stats.id = items.id
where
    items.type = 2
group by
    items.id,
    recipeId,
    stats.id,
    stats.durability,
    stats.gearLevel,
    stats.mainStat,
    stats.set`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Armour{}
		item.Stations = make([]int32, 0)
		item.ReagentFor = make([]int32, 0)
		item.Tags = make([]string, 0)
		item.BonusStats = make([]string, 0)
		item.Locations = make([]string, 0)

		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.Tier,
			pq.Array(&item.Stations),
			&item.Recipe,
			pq.Array(&item.ReagentFor),
			pq.Array(&item.Tags),
			&item.Durability,
			&item.GearLevel,
			&item.MainStat,
			&item.Set,
			pq.Array(&item.BonusStats),
			pq.Array(&item.Locations))

		if readError != nil {
			fmt.Println(readError)
			continue
		}
		items = append(items, item)
	}

	ctx.JSON(200, items)
}

func HandleArmourRequest(ctx *gin.RouterGroup) {
	ctx.GET("/list", getArmourList)
}
