package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func getRegionList(ctx *gin.Context) {
	var items = make([]models.Region, 0)
	connection := database.CreateConnection()
	defer connection.Close()
	query := `select * from regions`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Region{}
		readError := rows.Scan(&item.Id, &item.Name)
		if readError != nil {
			fmt.Print(readError)
			continue
		}
		items = append(items, item)
	}

	ctx.JSON(200, items)
}

func getLocationList(ctx *gin.Context) {

	var items = make([]models.Location, 0)
	connection := database.CreateConnection()
	defer connection.Close()
	query := `select
    locations.id,
    locations.name,
    locations.description,
    locations.mapgenieid,
    regions.name,
    array_agg(itemlocations.itemid)
from
    locations
    join regions on regions.id = locations.regionid
    join itemlocations on itemlocations.locationid = locations.id
group by
    locations.id,
    regions.name`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Location{}
		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.MapgenieId,
			&item.Region,
			pq.Array(&item.Loot))
		if readError != nil {
			fmt.Print(readError)
			continue
		}
		items = append(items, item)
	}

	ctx.JSON(200, items)
}

func HandleLocationRequest(ctx *gin.RouterGroup) {
	ctx.GET("/list", getLocationList)
	ctx.GET("/regions", getRegionList)
}
