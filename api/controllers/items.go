package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/database/tables"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ItemJson struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tier        int32  `json:"tier"`
}

func getItemsList(ctx *gin.Context) {

	var items = make([]ItemJson, 0)
	connection := database.CreateConnection()
	defer connection.Close()
	query := fmt.Sprintf(
		"SELECT %s, %s, %s FROM %s",
		tables.NameField,
		tables.DescriptionField,
		tables.ItemTierField,
		tables.ItemTableName)

	fmt.Println(query)
	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := ItemJson{}
		readError := rows.Scan(&item.Name, &item.Description, &item.Tier)
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
