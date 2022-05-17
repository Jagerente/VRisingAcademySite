package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func handleTagList(ctx *gin.Context) {
	var items = make([]models.Tag, 0)
	connection := database.CreateConnection()
	defer connection.Close()

	query := `select * from tags`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Tag{}

		readError := rows.Scan(
			&item.Id,
			&item.Value)

		if readError != nil {
			fmt.Println(readError)
			continue
		}

		items = append(items, item)
	}

	ctx.JSON(200, items)
}

func HandleTagRequest(group *gin.RouterGroup) {
	group.GET("/list", handleTagList)
}
