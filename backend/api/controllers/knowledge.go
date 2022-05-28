package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func getKnowledgeList(ctx *gin.Context) {

	var items = make([]models.Knowledge, 0)
	connection := database.CreateConnection()
	defer connection.Close()
	query := `select
	knowledges.id,
	array(
		(select recipes.id from recipes
		where recipes.knowledgeid=knowledges.id)
	) as recipes,
	array(
		(select spells.id from spells
		where spells.knowledgeid=knowledges.id)
	) as spells,
	quests.id as questId
	from knowledges
	join quests on quests.knowledgeid=knowledges.id
	group by knowledges.id, quests.id`

	fmt.Println(query)
	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		item := models.Knowledge{}
		readError := rows.Scan(
			&item.Id,
			pq.Array(&item.Recipes),
			pq.Array(&item.Spells),
			&item.Quest)

		if readError != nil {
			fmt.Print(readError)
			continue
		}
		items = append(items, item)
	}

	ctx.JSON(200, items)
}

func HandleKnowledgeRequest(ctx *gin.RouterGroup) {
	ctx.GET("/list", getKnowledgeList)
}
