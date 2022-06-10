package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func handleBloodTypeList(ctx *gin.Context) {
	var items = make([]models.BloodType, 0)
	connection := database.CreateConnection()
	defer connection.Close()

	query := `select * from bloodtypes`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.BloodType{
			Bonuses: make([]string, 0)}

		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			pq.Array(&item.Bonuses))

		if readError != nil {
			fmt.Println(readError)
			continue
		}

		items = append(items, item)
	}

	ctx.JSON(200, items)
}

func HandleBloodTypeRequest(group *gin.RouterGroup) {
	group.GET("/list", handleBloodTypeList)
}
