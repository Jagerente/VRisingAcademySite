package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type ItemJson struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tier        int32  `json:"tier"`
	Type        int32  `json:"type"`
}

type ListResponse struct {
}

func getItemTypesList() []models.ItemTypeObject {
	var items = make([]models.ItemTypeObject, 0)
	connection := database.CreateConnection()
	defer connection.Close()

	query := `select * from itemtypes order by itemtypes.id`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.ItemTypeObject{}

		readError := rows.Scan(
			&item.Id,
			&item.Name)

		if readError != nil {
			fmt.Println(readError)
			continue
		}

		items = append(items, item)
	}

	return items
}

func handleItemTypesList(ctx *gin.Context) {

	var items = make([]models.ItemTypeObject, 0)
	connection := database.CreateConnection()
	defer connection.Close()

	query := `select * from itemtypes`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.ItemTypeObject{}

		readError := rows.Scan(
			&item.Id,
			&item.Name)

		if readError != nil {
			fmt.Println(readError)
			continue
		}

		items = append(items, item)
	}
	ctx.JSON(200, items)
}

func handleListRequest(ctx *gin.Context) {
	var response = make([]models.Item, 0)
	connection := database.CreateConnection()
	defer connection.Close()

	newQuery := `select
    items.id,
    items.name,
    items.description,
    items.tier,
    items.type as typeid,
    itemtypes.title as typetitle,
    items.knowledgeid,
    array(
        (
            select
                stcns.id
            from
                recipestations as rcst
                join stations as stcns on rcst.stationId = stcns.id
            where
                rcst.recipeId = rcp.id
        )
    ) as stations,
    array(
        (
            select
                recipes.id
            from
                reciperesults
                join recipes on recipes.id = reciperesults.recipeid
            where
                reciperesults.itemid = items.id
        )
    ) as recipes,
    array(
        (
            select
                rc.id
            from
                recipeingredients as rci
                join recipes as rc on rci.recipeId = rc.id
            where
                rci.itemId = items.id
        )
    ) as reagentFor,
    array(
        (
            select
                tgs.value
            from
                itemtags as itt
                join tags as tgs on itt.tagId = tgs.id
            where
                itt.itemId = items.id
        )
    ) as tags,
    stats.durability,
    stats.gearLevel,
    stats.mainStat,
    items.setid,
    sets.name,
	sets.description,
    stats.slotid,
    array(
        (
            select
                secondarystats.bonus
            from
                secondaryitemstats
                join secondarystats on secondarystats.id = secondaryitemstats.secondarystatid
            where
                secondaryitemstats.statsId = stats.id
        )
    ) as bonusStats,
    array(
        (
            select
                lc.id
            from
                itemlocations as itl
                join locations as lc on itl.locationId = lc.id
            where
                itl.itemId = items.id
        )
    ) as locations,
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
    ) as variants
from
    items
    join itemtypes on itemtypes.id = items.type full
    join recipeingredients on recipeingredients.itemid = items.id full
    join recipes as rcp on recipeingredients.itemid = rcp.id full
    join itemstats as stats on stats.id = items.id full
    join sets on sets.id = items.setid
group by
    items.id,
    itemtypes.title,
    rcp.id,
    stats.id,
    stats.durability,
    stats.gearLevel,
    stats.mainStat,
    items.setid,
    sets.name,
    sets.description
order by
    items.id,
    items.type,
    items.setid`

	rows, err := connection.Query(newQuery)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		item := models.Item{
			Stations:   make([]int32, 0),
			Recipes:    make([]int32, 0),
			ReagentFor: make([]int32, 0),
			Tags:       make([]string, 0),
			BonusStats: make([]string, 0),
			Locations:  make([]int32, 0),
			Variants:   make([]int32, 0),
			Type:       models.ItemTypeObject{}}

		var setName *string
		var setDesc *string

		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.Tier,
			&item.Type.Id,
			&item.Type.Name,
			&item.KnowledgeId,
			pq.Array(&item.Stations),
			pq.Array(&item.Recipes),
			pq.Array(&item.ReagentFor),
			pq.Array(&item.Tags),
			&item.Durability,
			&item.GearLevel,
			&item.MainStat,
			&item.SetId,
			&setName,
			&setDesc,
			&item.SlotId,
			pq.Array(&item.BonusStats),
			pq.Array(&item.Locations),
			pq.Array(&item.Variants))

		if readError != nil {
			fmt.Print(readError)
			continue
		}
		if item.SetId != nil {
			item.Set = &models.ItemSetObject{
				Id:          *item.SetId,
				Name:        *setName,
				Description: setDesc}
		}
		response = append(response, item)
	}

	ctx.JSON(200, response)
}

func handleGroupedItemsListRequest(ctx *gin.Context) {
	type ResponseItem struct {
		Id    int32         `json:"id"`
		Name  string        `json:"name"`
		Items []models.Item `json:"items"`
	}
	type Response struct {
		Id   int32          `json:"id"`
		Name string         `json:"name"`
		Sets []ResponseItem `json:"sets"`
	}

	var response = make([]Response, 0)
	types := getItemTypesList()

	for _, item := range types {
		toAdd := Response{
			Id:   item.Id,
			Name: item.Name,
			Sets: make([]ResponseItem, 0)}
		response = append(response, toAdd)
	}

	type ResponseMapItem map[string][]models.Item
	type ResponseMap map[string]*ResponseMapItem
	const UnsetTypeId int32 = -1

	var responseMap = ResponseMap{}
	for _, item := range types {
		responseMap[item.Name] = &ResponseMapItem{}
	}

	var setMap = map[int32]string{}

	sets := getSetsList()
	for _, item := range sets {
		if _, ok := setMap[item.Id]; !ok {
			setMap[item.Id] = item.Name
		}
	}
	setMap[UnsetTypeId] = "Unset"

	connection := database.CreateConnection()
	defer connection.Close()

	newQuery := `select
    items.id,
    items.name,
    items.description,
    items.tier,
    items.type as typeid,
    itemtypes.title as typetitle,
    items.knowledgeid,
    array(
        (
            select
                distinct stcns.id
            from
                recipestations as rcst
                join stations as stcns on rcst.stationId = stcns.id
            where
                rcst.recipeId = any(array_agg((select recipeingredients.recipeId where recipeingredients.itemId=items.id)))
        )
    ) as stations,
    array(
        (
            select
                recipes.id
            from
                reciperesults
                join recipes on recipes.id = reciperesults.recipeid
            where
                reciperesults.itemid = items.id
        )
    ) as itemRecipes,
    array(
        (
            select
                rc.id
            from
                recipeingredients as rci
                join recipes as rc on rci.recipeId = rc.id
            where
                rci.itemId = items.id
        )
    ) as reagentFor,
    array(
        (
            select
                tgs.value
            from
                itemtags as itt
                join tags as tgs on itt.tagId = tgs.id
            where
                itt.itemId = items.id
        )
    ) as tags,
    stats.durability,
    stats.gearLevel,
    stats.mainStat,
    items.setid,
    sets.name,
	sets.description,
    stats.slotid,
    array(
        (
            select
                secondarystats.bonus
            from
                secondaryitemstats
                join secondarystats on secondarystats.id = secondaryitemstats.secondarystatid
            where
                secondaryitemstats.statsId = stats.id
        )
    ) as bonusStats,
    array(
        (
            select
                lc.id
            from
                itemlocations as itl
                join locations as lc on itl.locationId = lc.id
            where
                itl.itemId = items.id
        )
    ) as locations,
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
    ) as variants
from
    items
    join itemtypes on itemtypes.id = items.type 
    full join recipeingredients on recipeingredients.itemid = items.id
    full join recipes as rcp on recipeingredients.itemid = rcp.id
    full join itemstats as stats on stats.id = items.id
    full join sets on sets.id = items.setid
where items.id IS NOT NULL
group by
    items.id,
    itemtypes.title,
    rcp.id,
    stats.id,
    stats.durability,
    stats.gearLevel,
    stats.mainStat,
    items.setid,
    sets.name,
    sets.description
order by
    items.id,
    items.type,
    items.setid`

	rows, err := connection.Query(newQuery)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		item := models.Item{
			Stations:   make([]int32, 0),
			Recipes:    make([]int32, 0),
			ReagentFor: make([]int32, 0),
			Tags:       make([]string, 0),
			BonusStats: make([]string, 0),
			Locations:  make([]int32, 0),
			Variants:   make([]int32, 0),
			Type:       models.ItemTypeObject{}}

		var setName *string
		var setDesc *string

		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.Tier,
			&item.Type.Id,
			&item.Type.Name,
			&item.KnowledgeId,
			pq.Array(&item.Stations),
			pq.Array(&item.Recipes),
			pq.Array(&item.ReagentFor),
			pq.Array(&item.Tags),
			&item.Durability,
			&item.GearLevel,
			&item.MainStat,
			&item.SetId,
			&setName,
			&setDesc,
			&item.SlotId,
			pq.Array(&item.BonusStats),
			pq.Array(&item.Locations),
			pq.Array(&item.Variants))

		if readError != nil {
			fmt.Print(readError)
			continue
		}

		var key int32 = UnsetTypeId
		if item.SetId != nil {
			key = *item.SetId
			item.Set = &models.ItemSetObject{
				Id:          key,
				Name:        *setName,
				Description: setDesc}
		}

		var responseItem *ResponseMapItem = responseMap[item.Type.Name]
		if _, ok := (*responseItem)[setMap[key]]; ok {
			updatedArray := append((*responseItem)[setMap[key]], item)
			(*responseItem)[setMap[key]] = updatedArray
		} else {
			arayToAdd := []models.Item{item}
			(*responseItem)[setMap[key]] = arayToAdd
		}
	}

	var i int32 = 0
	for _, value := range responseMap {
		for _, valueValue := range *value {
			var newVal ResponseItem = ResponseItem{
				Id:   -1,
				Name: "Unset",
			}
			if valueValue[0].SetId != nil {
				newVal = ResponseItem{
					Id:    valueValue[0].Set.Id,
					Name:  valueValue[0].Set.Name,
					Items: make([]models.Item, 0)}
			}
			for _, itemValue := range valueValue {
				newVal.Items = append(newVal.Items, itemValue)
			}
			response[i].Sets = append(response[i].Sets, newVal)
		}
		i++
	}

	ctx.JSON(200, response)
}

func HandleItemsRequest(ctx *gin.RouterGroup) {
	ctx.GET("list", handleListRequest)
	ctx.GET("/grouplist", handleGroupedItemsListRequest)
	ctx.GET("/types", handleItemTypesList)
}
