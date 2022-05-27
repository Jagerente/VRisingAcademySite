package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func getQuestList(ctx *gin.Context) {

	var items = make([]models.Quest, 0)
	connection := database.CreateConnection()
	defer connection.Close()
	query := `select
	quests.id,
	quests.name,
	quests.description,
	quests.goals,
	quests.knowledgeid
	from quests`

	fmt.Println(query)
	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		item := models.Quest{}
		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			pq.Array(&item.Goals),
			&item.Knowledge)

		if readError != nil {
			fmt.Print(readError)
			continue
		}
		items = append(items, item)
	}

	ctx.JSON(200, items)
}

func HandleQuestRequest(ctx *gin.RouterGroup) {
	ctx.GET("/list", getQuestList)
}
