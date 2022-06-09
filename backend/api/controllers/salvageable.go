package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func handleSalvageableList(ctx *gin.Context) {
	var items = make([]models.Salvageable, 0)
	connection := database.CreateConnection()
	defer connection.Close()

	query := `select
    salvageables.id,
    salvageables.itemid,
    array_agg(salvageableresults.itemid) as itemids,
    array_agg(salvageableresults.amount) as itemamounts
from
    salvageables
    join salvageableresults on salvageableresults.salvageableid=salvageables.id
group by
    salvageables.id
order by
    salvageables.id`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Salvageable{
			Results: make([]models.SalvageableResult, 0)}
		ids := make([]int32, 0)
		amounts := make([]int32, 0)

		readError := rows.Scan(
			&item.Id,
			&item.ItemId,
			pq.Array(&ids),
			pq.Array(&amounts))

		if readError != nil {
			fmt.Println(readError)
			continue
		}

		for resultIndex := range ids {
			item.Results = append(item.Results, models.SalvageableResult{ItemId: ids[resultIndex], Amount: amounts[resultIndex]})
		}

		items = append(items, item)
	}

	ctx.JSON(200, items)
}

func HandleSalvageableRequest(group *gin.RouterGroup) {
	group.GET("/list", handleSalvageableList)
}
