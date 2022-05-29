package controllers

import (
	//"net/http"
	"VRisingAcademySite/database"
	"VRisingAcademySite/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SpellFilter struct {
	Type            int32
	School          int32
	IncludeSchoolId int32
}

func parseFilter(request *gin.Context) {

}

func handleSpellList(request *gin.Context) {
	filter := SpellFilter{
		Type:            0,
		School:          0,
		IncludeSchoolId: 0}
	if request.Query("school") != "" {
		parsedInt, err := strconv.Atoi(request.Query("school"))
		if err == nil {
			filter.School = int32(parsedInt)
		}
	}
	if request.Query("type") != "" {
		parsedInt, err := strconv.Atoi(request.Query("type"))
		if err == nil {
			filter.Type = int32(parsedInt)
		}
	}
	if request.Query("includeSchoolId") != "" {
		parsedInt, err := strconv.Atoi(request.Query("includeSchoolId"))
		if err == nil {
			filter.IncludeSchoolId = int32(parsedInt)
		}
	}
	connection := database.CreateConnection()
	defer connection.Close()

	query := `select
    spells.id,
    spells.name,
    spells.description,
	spellschools.id as spellschoolid,
    spellschools.name,
    spelltypes.title,
    spells.cooldown,
    spells.casttime,
    spells.charges,
	spells.knowledgeId
from spells
join spellschools on spellschools.id=spells.schoolid
join spelltypes on spelltypes.id=spells.typeid`

	if filter.School > 0 {
		query += fmt.Sprintf(" where spellschools.id=%d", filter.School)
	}
	if filter.Type > 0 {
		query += fmt.Sprintf(" where spelltypes.id=%d", filter.Type)
	}

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		fmt.Println(query)
		panic(err)
	}
	defer rows.Close()

	if filter.IncludeSchoolId > 0 {

		var items = make([]models.SpellWithSchoolId, 0)
		for rows.Next() {
			item := models.SpellWithSchoolId{}

			readError := rows.Scan(
				&item.Id,
				&item.Name,
				&item.Description,
				&item.SchoolId,
				&item.School,
				&item.Type,
				&item.Cooldown,
				&item.CastTime,
				&item.Charges,
				&item.Knowledge)

			if readError != nil {
				fmt.Println(readError)
				continue
			}

			items = append(items, item)
		}

		request.JSON(200, items)
	} else {

		var items = make([]models.Spell, 0)
		for rows.Next() {
			item := models.Spell{}
			var id int32

			readError := rows.Scan(
				&item.Id,
				&item.Name,
				&item.Description,
				&id,
				&item.School,
				&item.Type,
				&item.Cooldown,
				&item.CastTime,
				&item.Charges,
				&item.Knowledge)

			if readError != nil {
				fmt.Println(readError)
				continue
			}

			items = append(items, item)
		}

		request.JSON(200, items)
	}
}

func HandleSpellRequest(group *gin.RouterGroup) {
	group.GET("/list", handleSpellList)
}
