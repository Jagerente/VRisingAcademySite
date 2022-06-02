package database

import (
	"VRisingAcademySite/database/tables"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	connectionString         string = "user=%s password=%s host=%s port=%d dbname=vrisingdb sslmode=disable"
	connectionStringBaseless string = "user=%s password=%s host=%s port=%d sslmode=disable"
	checkConnectionString    string = `SELECT * FROM CheckTable limit 1`
	postgresUser             string = "postgres"
	postgresPassword         string = "83"
	postgresHost             string = "localhost"
	postgresPort             int32  = 5432
	DatabaseVersion          int32  = 7
)

func CheckIfDatabaseNeedsUpdate() bool {
	connection := CreateConnection()
	var id, exists, version int32
	item := connection.QueryRow(checkConnectionString)
	readError := item.Scan(&id, &exists, &version)

	if readError != nil {
		panic(readError)
	}

	return version < DatabaseVersion
}

func UpdateDatabase() {
	connection := CreateConnection()
	defer connection.Close()
	var id, exists, version int32
	item := connection.QueryRow(checkConnectionString)
	readError := item.Scan(&id, &exists, &version)

	if readError != nil {
		panic(readError)
	}

	if version >= DatabaseVersion {
		return
	}

	startIndex := version - 1

	for index := startIndex; index < DatabaseVersion-1; index++ {
		for _, migrationQuery := range migrations[index] {
			connection.Exec(migrationQuery)
		}
	}

	connection.Exec(fmt.Sprintf(`UPDATE CheckTable SET Version=%d`, DatabaseVersion))
}

//Checks if database exists
func CheckIfDatabaseExists() bool {
	connection, _ := sql.Open("postgres", fmt.Sprintf(connectionString,
		postgresUser,
		postgresPassword,
		postgresHost,
		postgresPort))

	var id, result1, result2 int32
	item := connection.QueryRow(checkConnectionString)
	readError := item.Scan(&id, &result1, &result2)
	connection.Close()

	if readError != nil {
		return false
	}

	return true
}

//Creates connection to database
func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf(connectionString,
		postgresUser,
		postgresPassword,
		postgresHost,
		postgresPort))
	if err != nil {
		panic(err)
	}
	return db
}

//Initializes database
func InitializeDatabase() {
	connection, err := sql.Open("postgres", fmt.Sprintf(connectionStringBaseless,
		postgresUser,
		postgresPassword,
		postgresHost,
		postgresPort))
	if err != nil {
		panic(err)
	}

	_, k := connection.Query("CREATE DATABASE vrisingdb;")
	if k != nil {
		connection.Close()
		panic(k)
	}
	connection.Close()

	connection = CreateConnection()
	defer connection.Close()

	connection.Exec("CREATE TABLE CheckTable (Id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY, Exists INTEGER NOT NULL, Version INTEGER NOT NULL);")

	//Основные сущности
	//Предметы
	connection.Exec(fmt.Sprintf(
		tables.BaseCreateGameEntityQuery,
		tables.ItemTableName,
		fmt.Sprintf(`%s INTEGER NOT NULL DEFAULT 0,
		type INTEGER NOT NULL DEFAULT 0`, tables.ItemTierField)))

	//Рецепты
	connection.Exec(fmt.Sprintf(
		tables.BaseCreateEntityQuery,
		tables.RecipeTableName,
		fmt.Sprintf(`time INTEGER NOT NULL DEFAULT 60`)))

	//Тэги предметов
	connection.Exec(fmt.Sprintf(
		tables.BaseCreateEntityQuery,
		tables.TagsTableName,
		fmt.Sprintf("%s varchar(64) NOT NULL UNIQUE", tables.TagValueField)))

	//Монстры
	connection.Exec(fmt.Sprintf(
		tables.BaseCreateGameEntityQuery,
		tables.MonsterTableName,
		fmt.Sprintf(`%s SMALLINT NOT NULL DEFAULT 0`, tables.IsBossFieldName)))

	//Локации
	connection.Exec(fmt.Sprintf(
		`CREATE TABLE %s (
			id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
			value VARCHAR(256)
		);`,
		tables.LocationTableName))

	connection.Exec(
		`CREATE TABLE sets (
			id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
			name varchar(128) NOT NULL,
			description varchar(512) NOT NULL
		);`)

	//Статы предметов
	connection.Exec(
		`CREATE TABLE itemstats (
			id INTEGER PRIMARY KEY NOT NULL REFERENCES items(id),
			mainStat DOUBLE PRECISION,
			setId INTEGER REFERENCES sets(id),
			setBonus VARCHAR(256),
			gearLevel INTEGER,
			durability INTEGER,
			slotId INTEGER
		);`)

	//Таблицы для связей многие-ко-многим
	//Таблица для связи многие ко многим между предметами и тегами
	connection.Exec(fmt.Sprintf(
		`CREATE TABLE %s (
        %s INTEGER NOT NULL REFERENCES %s(%s) ON DELETE CASCADE,
        %s INTEGER NOT NULL REFERENCES %s(%s) ON DELETE CASCADE,
        PRIMARY KEY(%s, %s)
        );`,
		tables.ItemTagTableName,
		tables.ItemIdFieldName, tables.ItemTableName, tables.IdFieldName,
		tables.TagIdFieldName, tables.TagsTableName, tables.IdFieldName,
		tables.ItemIdFieldName, tables.TagIdFieldName))

	//Таблица для связи многие ко многим между монстрами и предмтами (дроп с монстров)
	connection.Exec(fmt.Sprintf(
		`CREATE TABLE %s (
		%s INTEGER NOT NULL REFERENCES %s(%s) ON DELETE CASCADE,
		%s INTEGER NOT NULL REFERENCES %s(%s) ON DELETE CASCADE,
		PRIMARY KEY(%s, %s)
		);`,
		tables.MonsterDropTableName,
		tables.MonsterIdFieldName, tables.MonsterTableName, tables.IdFieldName,
		tables.ItemIdFieldName, tables.ItemTableName, tables.IdFieldName,
		tables.MonsterIdFieldName, tables.ItemIdFieldName))

	//Таблица для связи многие ко многим между монстрами и локациями
	connection.Exec(fmt.Sprintf(
		`CREATE TABLE %s (
		%s INTEGER NOT NULL REFERENCES %s(%s) ON DELETE CASCADE,
		%s INTEGER NOT NULL REFERENCES %s(%s) ON DELETE CASCADE,
		PRIMARY KEY(%s, %s)
		);`,
		tables.MonsterLocationTableName,
		tables.MonsterIdFieldName, tables.MonsterTableName, tables.IdFieldName,
		tables.LocationIdFieldName, tables.LocationTableName, tables.IdFieldName,
		tables.MonsterIdFieldName, tables.LocationIdFieldName))

	//Таблица для связи многие ко многим между рецептами и предметами (их ингридиентами)
	connection.Exec(fmt.Sprintf(
		`CREATE TABLE %s (
		%s INTEGER NOT NULL REFERENCES %s(%s) ON DELETE CASCADE,
		%s INTEGER NOT NULL REFERENCES %s(%s) ON DELETE CASCADE,
		amount INTEGER NOT NULL DEFAULT 1,
		PRIMARY KEY(%s, %s)
		);`,
		tables.RecipeItemTableName,
		tables.RecipeIdFieldName, tables.RecipeTableName, tables.IdFieldName,
		tables.ItemIdFieldName, tables.ItemTableName, tables.IdFieldName,
		tables.RecipeIdFieldName, tables.ItemIdFieldName))

	//Локации предметов
	connection.Exec(
		`CREATE TABLE itemlocations (
			itemId INTEGER NOT NULL REFERENCES items(id) ON DELETE CASCADE,
			locationId INTEGER NOT NULL REFERENCES locations(id) ON DELETE CASCADE,
			PRIMARY KEY(itemId, locationId)
		);`)

	connection.Exec(
		`CREATE TABLE secondarystats (
			id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
			bonus VARCHAR(512) NOT NULL
		);`)

	connection.Exec(
		`CREATE TABLE secondaryitemstats (
			statsId INTEGER NOT NULL REFERENCES itemstats(id) ON DELETE CASCADE,
			secondaryStatId INTEGER NOT NULL REFERENCES secondarystats(id) ON DELETE CASCADE
		);`)

	connection.Exec(
		`CREATE TABLE stations(
			id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
			name VARCHAR(128) NOT NULL,
			description VARCHAR(512)
		);`)

	connection.Exec(
		`CREATE TABLE recipestations (
			recipeId INTEGER NOT NULL REFERENCES recipes(id) ON DELETE CASCADE,
			stationId INTEGER NOT NULL REFERENCES stations(id) ON DELETE CASCADE
		);`)

	connection.Exec(fmt.Sprintf(
		`CREATE TABLE %s (
		%s INTEGER NOT NULL REFERENCES %s(%s) ON DELETE CASCADE,
		%s INTEGER NOT NULL REFERENCES %s(%s) ON DELETE CASCADE,
		amount INTEGER NOT NULL DEFAULT 1,
		PRIMARY KEY(%s, %s)
		);`,
		"reciperesults",
		tables.RecipeIdFieldName, tables.RecipeTableName, tables.IdFieldName,
		tables.ItemIdFieldName, tables.ItemTableName, tables.IdFieldName,
		tables.RecipeIdFieldName, tables.ItemIdFieldName))

	//Заполнение поля о существовании таблицы
	connection.Exec("INSERT INTO CheckTable(Exists, Version) VALUES (1, 1);")

	connection.Close()

	UpdateDatabase()
}
