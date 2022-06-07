package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func getSetsList() []models.Set {
	var items = make([]models.Set, 0)
	connection := database.CreateConnection()
	defer connection.Close()

	query := `select * from sets`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Set{}

		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description)

		if readError != nil {
			fmt.Println(readError)
			continue
		}

		items = append(items, item)
	}
	return items
}

func handleSetList(ctx *gin.Context) {
	items := getSetsList()
	ctx.JSON(200, items)
}

func HandleSetRequest(group *gin.RouterGroup) {
	group.GET("/list", handleSetList)
}
