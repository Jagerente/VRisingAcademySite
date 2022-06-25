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
	LocationQuery = `select
    locations.id,
    locations.name,
    locations.description,
    locations.mapgenieid,
    regions.name,
    array_remove(array_agg(itemlocations.itemid), NULL) as items,
	array_remove(array_agg(locationshazards.hazardid), NULL) as hazardIds
from
    locations
    left join regions on regions.id = locations.regionid
    left join itemlocations on itemlocations.locationid = locations.id
	left join locationshazards on locationshazards.locationid = locations.id
group by
    locations.id,
    regions.name`
	HazardQuery = `select
	hazards.id,
	hazards.name,
	hazards.description,
	array_remove(array_agg(locationshazards.locationid), NULL) as locations
from
    hazards
	left join locationshazards on locationshazards.hazardid = hazards.id
group by
    hazards.id`
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

	rows, err := connection.Query(LocationQuery)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Location{
			Loot:    make([]int32, 0),
			Hazards: make([]int32, 0),
		}
		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.MapgenieId,
			&item.Region,
			pq.Array(&item.Loot),
			pq.Array(&item.Hazards))
		if readError != nil {
			fmt.Print(readError)
			continue
		}
		items = append(items, item)
	}

	ctx.JSON(200, items)
}

func getHazardList(ctx *gin.Context) {

	var items = make([]models.Hazard, 0)
	connection := database.CreateConnection()
	defer connection.Close()

	rows, err := connection.Query(HazardQuery)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Hazard{
			Locations: make([]int32, 0)}
		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			pq.Array(&item.Locations))
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
	ctx.GET("/hazards", getHazardList)
}
