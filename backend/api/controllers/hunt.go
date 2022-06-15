package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func getHuntList() []models.Hunt {
	var items = make([]models.Hunt, 0)
	connection := database.CreateConnection()
	defer connection.Close()

	query := `select
    hunts.id,
    regions.name as region,
    hunts.location,
    hunts.maxservants,
    hunts.difficulty,
    hunts.classbonus,
    array_agg(huntrewardvalues.itemid) as itemids,
    array_agg(huntmodifiers.servantsamount) as servants,
    array_agg(huntmodifiers.bonus) as bonus,
    array_agg(hunttimes.time) as time,
    array_agg(huntrewardvalues.minamount) as minimalamounts,
    array_agg(huntrewardvalues.maxamount) as maximumamounts
from
    hunts
    join regions on regions.id = hunts.regionid
    join huntrewardvalues on huntrewardvalues.huntid = hunts.id
    join huntmodifiers on huntmodifiers.id = huntrewardvalues.modifier
    join hunttimes on hunttimes.id = huntrewardvalues.time
group by
    hunts.id,
    regions.name
order by
    hunts.id`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Hunt{
			Rewards: make([]models.HuntRewardValue, 0)}

		itemids := make([]int32, 0)
		servants := make([]int32, 0)
		bonuses := make([]string, 0)
		times := make([]int32, 0)
		minimals := make([]int32, 0)
		maximums := make([]int32, 0)

		readError := rows.Scan(
			&item.Id,
			&item.Region,
			&item.Location,
			&item.MaxServants,
			&item.Difficulty,
			&item.BonusId,
			pq.Array(&itemids),
			pq.Array(&servants),
			pq.Array(&bonuses),
			pq.Array(&times),
			pq.Array(&minimals),
			pq.Array(&maximums))

		if readError != nil {
			fmt.Println(readError)
			continue
		}

		for index := range itemids {
			item.Rewards = append(item.Rewards, models.HuntRewardValue{
				ItemId:        itemids[index],
				Servants:      servants[index],
				Bonus:         bonuses[index],
				Time:          times[index],
				MinimalAmount: minimals[index],
				MaximumAmount: maximums[index]})
		}

		items = append(items, item)
	}
	return items
}

func handleHuntList(ctx *gin.Context) {
	items := getHuntList()
	ctx.JSON(200, items)
}

func HandleHuntRequest(group *gin.RouterGroup) {
	group.GET("/list", handleHuntList)
}
