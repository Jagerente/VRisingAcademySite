package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func handleStationList(ctx *gin.Context) {
	var items = make([]models.Station, 0)
	connection := database.CreateConnection()
	defer connection.Close()

	query := `select
    stations.id,
    stations.name,
    stations.description,
    array(
        (
            with rctst as (
                select
                    recipes.id as recipeid
                from
                    recipestations
                    join recipes on recipes.id = recipestations.recipeid
                where
                    recipestations.stationid = stations.id
            )
            select
                reciperesults.itemid
            from
                reciperesults
                join rctst on reciperesults.recipeid = rctst.recipeid
        )
    ) as recipes,
    array(
        (
            select
                recipes.id
            from
                recipestations
                join recipes on recipes.id = recipestations.recipeid
            where
                recipestations.stationid = stations.id
        )
    ) as stationRecipes,
    array(
        (
            select
                recipes.knowledgeid
            from
                recipestations
                join recipes on recipes.id = recipestations.recipeid
            where
                recipestations.stationid = stations.id
        )
    ) as knowledges
from
    stations`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Station{}
		item.Recipes = make([]int32, 0)
		item.StationRecipes = make([]int32, 0)
		item.Knowledges = make([]int32, 0)

		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			pq.Array(&item.Recipes),
			pq.Array(&item.StationRecipes),
			pq.Array(&item.Knowledges))

		if readError != nil {
			fmt.Println(readError)
			continue
		}

		items = append(items, item)
	}

	ctx.JSON(200, items)
}

func HandleStationRequest(group *gin.RouterGroup) {
	group.GET("/list", handleStationList)
}
