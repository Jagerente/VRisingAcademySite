package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func handleRecipeList(ctx *gin.Context) {
	var items = make([]models.Recipe, 0)
	connection := database.CreateConnection()
	defer connection.Close()

	query := `select
    recipes.id,
    recipes.name,
    recipes.description,
    recipes.resultitemid,
    recipes.time,
    array(
            (
                select
                    recipeingredients.itemid
                from
                    recipeingredients
                where
                    recipeingredients.recipeid = recipes.id
            )
    ) as ingredientsIds,
    array(
            (
                select
                    recipeingredients.amount
                from
                    recipeingredients
                where
                    recipeingredients.recipeid = recipes.id
            )
    ) as ingredientsAmounts
from
    recipes
group by
    recipes.id`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Recipe{}
		item.Ingredients = make([]models.RecipeIngredient, 0)
		ids := make([]int32, 0)
		amounts := make([]int32, 0)

		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.Result,
			&item.Time,
			pq.Array(&ids),
			pq.Array(&amounts))

		if readError != nil {
			fmt.Println(readError)
			continue
		}

		for index, _ := range ids {
			item.Ingredients = append(item.Ingredients, models.RecipeIngredient{Id: ids[index], Amount: amounts[index]})
		}

		items = append(items, item)
	}

	ctx.JSON(200, items)
}

func HandleRecipeRequest(group *gin.RouterGroup) {
	group.GET("/list", handleRecipeList)
}
