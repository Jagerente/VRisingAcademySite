package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func getReagentList(ctx *gin.Context) {

	var items = make([]models.Reagent, 0)
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
    array(
        (select recipes.id from reciperesults
        join recipes on recipes.id=reciperesults.recipeid
        where reciperesults.itemid=items.id)
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
    join recipeingredients on recipeingredients.itemid = items.id
    full join recipes as rcp on recipeingredients.itemid = rcp.id
where
    items.type = 4
group by
    items.id,
    rcp.id`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Reagent{}
		item.Stations = make([]int32, 0)
		item.Recipes = make([]int32, 0)
		item.ReagentFor = make([]int32, 0)
		item.Tags = make([]string, 0)
		item.Locations = make([]string, 0)

		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.Tier,
			pq.Array(&item.Stations),
			pq.Array(&item.Recipes),
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

func HandleReagentRequest(ctx *gin.RouterGroup) {
	ctx.GET("/list", getReagentList)
}
