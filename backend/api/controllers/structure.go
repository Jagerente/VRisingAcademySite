package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func handleStructureList(ctx *gin.Context) {
	var items = make([]models.Structure, 0)
	connection := database.CreateConnection()
	defer connection.Close()

	query := `select
    items.id,
    items.name,
    items.description,
    items.tier,
    items.type,
    array(
        (
            select
                structurevariants.variantid
            from
                structurevariants
            where
                structurevariants.structureid = items.id
			order by
			    structurevariants.variantid
        )
    ) as variants,
    array(
        (
            select
                reciperesults.recipeid
            from
                reciperesults
            where
                reciperesults.itemid = items.id
        )
    ) as recipes
from
    items
where
    items.type = 5
order by
    items.id`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Structure{
			IsVariant: false,
			Variants:  make([]int32, 0),
			Recipes:   make([]int32, 0)}

		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.Tier,
			&item.Type,
			pq.Array(&item.Variants),
			pq.Array(&item.Recipes))

		if readError != nil {
			fmt.Println(readError)
			continue
		}

		items = append(items, item)
	}

	if len(items) != 0 {
		var checkItem *models.Structure = &items[0]
		for i := 1; i < len(items); i++ {
			var item *models.Structure = &items[i]
			if len(item.Variants) == 0 || item.Variants[0] != checkItem.Id {
				checkItem = item
				continue
			} else {
				item.IsVariant = true
			}

		}
	}

	ctx.JSON(200, items)
}

func HandleStructureRequest(group *gin.RouterGroup) {
	group.GET("/list", handleStructureList)
}
