package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ItemJson struct {
	id          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tier        int32  `json:"tier"`
	Type        int32  `json:"type"`
}

func getItemsList(ctx *gin.Context) {

	var items = make([]ItemJson, 0)
	connection := database.CreateConnection()
	defer connection.Close()
	query := "select * from items"

	fmt.Println(query)
	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := ItemJson{}
		readError := rows.Scan(&item.id, &item.Name, &item.Description, &item.Tier, &item.Type)
		if readError != nil {
			fmt.Print(readError)
			continue
		}
		items = append(items, item)
	}

	ctx.JSON(200, items)
}

func HandleItemsRequest(ctx *gin.RouterGroup) {
	ctx.GET("/list", getItemsList)
}
