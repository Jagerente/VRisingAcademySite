package api

import (
	"VRisingAcademySite/api/controllers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BasicQueryFilter struct {
	Amount            int32
	Offset            int32
	NameFilter        string
	DescriptionFilter string
}

func ParseQueryFilter(request *gin.Context) BasicQueryFilter {
	result := BasicQueryFilter{
		Amount:            10,
		Offset:            0,
		NameFilter:        "",
		DescriptionFilter: ""}

	//Amount query parameter
	amountQuery := request.Query("amount")
	if amountQuery != "" {
		parsedInt, err := strconv.Atoi(amountQuery)
		if err == nil {
			if parsedInt >= 10 {
				result.Amount = int32(parsedInt)
			}
		}
	}

	//Offset query parameter
	offsetQuery := request.Query("offset")
	if offsetQuery != "" {
		parsedInt, err := strconv.Atoi(offsetQuery)
		if err == nil {
			if parsedInt >= 0 {
				result.Offset = int32(parsedInt)
			}
		}
	}

	//Filter by name parameter
	if request.Query("nameFilter") != "" {
		result.NameFilter = request.Query("nameFilter")
	}

	//Filter by description parameter
	if request.Query("descFilter") != "" {
		result.DescriptionFilter = request.Query("descFilter")
	}

	return result
}

func RegisterApiHandlers(r *gin.Engine) {
	items := r.Group("/api/items")
	{
		controllers.HandleItemsRequest(items)
	}

	weapons := r.Group("/api/weapon")
	{
		controllers.HandleWeaponsRequest(weapons)
	}
	armours := r.Group("/api/armour")
	{
		controllers.HandleArmourRequest(armours)
	}
	consumables := r.Group("/api/consumable")
	{
		controllers.HandleConsumableRequest(consumables)
	}
	reagents := r.Group("/api/reagent")
	{
		controllers.HandleReagentRequest(reagents)
	}
	recipes := r.Group("api/recipe")
	{
		controllers.HandleRecipeRequest(recipes)
	}
	tags := r.Group("/api/tag")
	{
		controllers.HandleTagRequest(tags)
	}
	sets := r.Group("api/set")
	{
		controllers.HandleSetRequest(sets)
	}
	stations := r.Group("/api/station")
	{
		controllers.HandleStationRequest(stations)
	}
	spells := r.Group("/api/spell")
	{
		controllers.HandleSpellRequest(spells)
	}
}
