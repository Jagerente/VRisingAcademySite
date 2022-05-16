package database

var (
	migrations [][]string = [][]string{
		[]string{
			`CREATE TABLE spellschools (
				id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	            name VARCHAR(128) NOT NULL,
	            description VARCHAR(512));`,
			`CREATE TABLE spelltypes (
				id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
				title VARCHAR(128) NOT NULL);`,
			`CREATE TABLE spells (
				id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	            name VARCHAR(128) NOT NULL,
	            description VARCHAR(512),
				schoolId INTEGER NOT NULL REFERENCES spellschools(id) ON DELETE SET NULL,
				typeId INTEGER NOT NULL REFERENCES spelltypes(id) ON DELETE SET NULL,
				cooldown DOUBLE PRECISION NOT NULL,
				casttime DOUBLE PRECISION NOT NULL,
				charges INTEGER NOT NULL DEFAULT 1);`}}
)
