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

func parseFilter(request *gin.Context) SpellFilter {
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
	return filter
}

func getSpellTypesList() []models.SpellType {
	var items = make([]models.SpellType, 0)
	connection := database.CreateConnection()
	defer connection.Close()
	query := "select * from spelltypes"

	fmt.Println(query)
	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.SpellType{}
		readError := rows.Scan(&item.Id, &item.Name)
		if readError != nil {
			fmt.Print(readError)
			continue
		}
		items = append(items, item)
	}
	return items
}

func getSpellSchoolsList() []models.SpellSchool {
	var items = make([]models.SpellSchool, 0)
	connection := database.CreateConnection()
	defer connection.Close()
	query := "select * from spellschools"

	fmt.Println(query)
	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.SpellSchool{}
		readError := rows.Scan(&item.Id, &item.Name, &item.Description)
		if readError != nil {
			fmt.Print(readError)
			continue
		}
		items = append(items, item)
	}
	return items
}

type SpellGroupedListResponseItem struct {
	Id     int32          `json:"id"`
	Name   string         `json:"name"`
	Spells []models.Spell `json:"spells"`
}
type SpellGroupedListResponse struct {
	Id    int32                          `json:"id"`
	Name  string                         `json:"name"`
	Types []SpellGroupedListResponseItem `json:"types"`
}

func findSpellTypeIndex(sets []SpellGroupedListResponseItem, spellType int32) int32 {
	for index, item := range sets {
		if item.Id == spellType {
			return int32(index)
		}
	}
	return -2
}

func findSchoolIndex(response []SpellGroupedListResponse, spellSchool int32) int32 {
	for index, item := range response {
		if item.Id == spellSchool {
			return int32(index)
		}
	}
	return -1
}

func handleGroupedList(request *gin.Context) {
	var response = make([]SpellGroupedListResponse, 0)
	schools := getSpellSchoolsList()

	for _, item := range schools {
		toAdd := SpellGroupedListResponse{
			Id:    item.Id,
			Name:  item.Name,
			Types: make([]SpellGroupedListResponseItem, 0)}
		response = append(response, toAdd)
	}

	connection := database.CreateConnection()
	defer connection.Close()

	query := `select
    spells.id,
    spells.name,
    spells.description,
	spellschools.id as spellschoolid,
    spellschools.name,
	spells.typeid,
    spelltypes.title,
    spells.cooldown,
    spells.casttime,
    spells.charges,
	spells.knowledgeId
from spells
join spellschools on spellschools.id=spells.schoolid
join spelltypes on spelltypes.id=spells.typeid
order by spellschools.id`

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		fmt.Println(query)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Spell{}

		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.SchoolId,
			&item.SchoolName,
			&item.TypeId,
			&item.TypeName,
			&item.Cooldown,
			&item.CastTime,
			&item.Charges,
			&item.Knowledge)

		if readError != nil {
			fmt.Println(readError)
			continue
		}

		item.School = models.SpellSchoolObject{
			Id:   item.SchoolId,
			Name: item.SchoolName}
		item.Type = models.SpellTypeObject{
			Id:   item.TypeId,
			Name: item.TypeName}

		schoolIndex := findSchoolIndex(response, item.School.Id)
		responseType := response[schoolIndex]
		typeIndex := findSpellTypeIndex(responseType.Types, item.Type.Id)
		if typeIndex == -2 {
			var newItem SpellGroupedListResponseItem = SpellGroupedListResponseItem{
				Id:     item.Type.Id,
				Name:   item.Type.Name,
				Spells: make([]models.Spell, 0)}
			responseType.Types = append(responseType.Types, newItem)
			typeIndex = findSpellTypeIndex(responseType.Types, item.Type.Id)
		}
		responseType.Types[typeIndex].Spells = append(responseType.Types[typeIndex].Spells, item)
		response[schoolIndex] = responseType
	}

	request.JSON(200, response)
}

func handleSchoolList(request *gin.Context) {
	items := getSpellSchoolsList()
	request.JSON(200, items)
}

func handleTypeRequest(request *gin.Context) {
	items := getSpellTypesList()
	request.JSON(200, items)
}

func handleSpellList(request *gin.Context) {
	filter := parseFilter(request)
	connection := database.CreateConnection()
	defer connection.Close()

	query := `select
    spells.id,
    spells.name,
    spells.description,
	spellschools.id as spellschoolid,
    spellschools.name,
	spells.typeid,
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
		if filter.School > 0 {
			query += " and"
		}
		query += fmt.Sprintf(" where spelltypes.id=%d", filter.Type)
	}

	rows, err := connection.Query(query)

	if err != nil {
		fmt.Print(err)
		fmt.Println(query)
		panic(err)
	}
	defer rows.Close()

	var items = make([]models.Spell, 0)
	for rows.Next() {
		item := models.Spell{}

		readError := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Description,
			&item.SchoolId,
			&item.SchoolName,
			&item.TypeId,
			&item.TypeName,
			&item.Cooldown,
			&item.CastTime,
			&item.Charges,
			&item.Knowledge)

		if readError != nil {
			fmt.Println(readError)
			continue
		}
		item.School = models.SpellSchoolObject{
			Id:   item.SchoolId,
			Name: item.SchoolName}
		item.Type = models.SpellTypeObject{
			Id:   item.TypeId,
			Name: item.TypeName}

		items = append(items, item)
	}

	request.JSON(200, items)
}

func HandleSpellRequest(group *gin.RouterGroup) {
	group.GET("/list", handleSpellList)
	group.GET("/schools", handleSchoolList)
	group.GET("/types", handleTypeRequest)
	group.GET("/grouplist", handleGroupedList)
}
