package database

var (
	migrations [][]string = [][]string{
		[]string{ //v2
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
				charges INTEGER NOT NULL DEFAULT 1);`},
		[]string{ //v3
			`CREATE TABLE knowledges (id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY);`,
			`CREATE TABLE quests (
				id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
				name VARCHAR(128) NOT NULL,
				description VARCHAR(512) NOT NULL,
				goals text[] NOT NULL,
				knowledgeId INTEGER NOT NULL REFERENCES knowledges(id));`,
			`CREATE TABLE stationknowledges (
				stationId INTEGER NOT NULL REFERENCES stations(id) ON DELETE CASCADE,
				knowledgeId INTEGER NOT NULL REFERENCES knowledges(id) ON DELETE CASCADE);`,
			`ALTER TABLE spells
			ADD knowledgeId INTEGER NULL REFERENCES knowledges(id) ON DELETE CASCADE;`,
			`ALTER TABLE recipes
			ADD knowledgeId INTEGER NULL REFERENCES knowledges(id) ON DELETE CASCADE;`},
		[]string{ //v4
			`CREATE TABLE structurevariants (
				structureid INTEGER NOT NULL REFERENCES items(id) ON DELETE CASCADE,
				variantid INTEGER NOT NULL REFERENCES items(id) ON DELETE CASCADE);`},
		[]string{ //v5
			`CREATE TABLE regions(id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY, name varchar(64) NOT NULL);`,
			`ALTER TABLE locations
			ADD COLUMN description varchar(512) NOT NULL DEFAULT ''`,
			`ALTER TABLE locations
			ADD COLUMN regionid INTEGER REFERENCES regions(id) ON DELETE CASCADE NULL`,
			`ALTER TABLE locations
			ADD COLUMN mapgenieid INTEGER NOT NULL DEFAULT 0`,
			`ALTER TABLE locations
			RENAME COLUMN value TO name`},
		[]string{ //v6
			`CREATE TABLE itemtypes (id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY, title varchar(64) NOT NULL);`,
			`ALTER TABLE items
			ADD COLUMN knowledgeid INTEGER REFERENCES knowledges(id) DEFAULT NULL;`,
			`ALTER TABLE items
			ADD FOREIGN KEY (type) REFERENCES itemtypes(id) NOT VALID;`},
		[]string{ //v7
			`ALTER TABLE items
			ADD COLUMN setid INTEGER NULL REFERENCES sets(id);`,
			`UPDATE items
			SET setid=itemstats.setid
			FROM itemstats WHERE itemstats.id=items.id`,
			`ALTER TABLE itemstats
			DROP COLUMN setid`},
		[]string{ //v8
			`ALTER TABLE items
			ADD COLUMN ownerid INTEGER REFERENCES items(id) DEFAULT NULL;`},
		[]string{ //v9
			`ALTER TABLE items
			DROP COLUMN ownerid;`},
		[]string{ //v10
			`CREATE TABLE salvageables(
				id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
				itemid INTEGER NOT NULL REFERENCES items(id));`,
			`CREATE TABLE salvageableresults(
				itemId INTEGER NOT NULL REFERENCES items(id) ON DELETE CASCADE,
				salvageableid INTEGER NOT NULL REFERENCES salvageables(id) ON DELETE CASCADE,
				amount INTEGER NOT NULL DEFAULT 1,
				PRIMARY KEY(itemId, salvageableid));`},
		[]string{
			`CREATE TABLE bloodtypes(
				id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
				name varchar(128) NOT NULL,
				description varchar(512) NOT NULL DEFAULT '',
				bonuses text [] DEFAULT NULL
			);`}}
)
