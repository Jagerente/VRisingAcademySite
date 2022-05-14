package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func getConsumableList(ctx *gin.Context) {

	var items = make([]models.Consumable, 0)
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
where
    items.type = 3
group by
    items.id,
    recipeId`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Consumable{}
		item.Stations = make([]int32, 0)
		item.ReagentFor = make([]int32, 0)
		item.Tags = make([]string, 0)
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
			pq.Array(&item.Locations))

		if readError != nil {
			fmt.Println(readError)
			continue
		}
		items = append(items, item)
	}

	ctx.JSON(200, items)
}

func HandleConsumableRequest(ctx *gin.RouterGroup) {
	ctx.GET("/list", getConsumableList)
}
