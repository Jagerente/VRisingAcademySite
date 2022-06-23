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
		[]string{ //v11
			`CREATE TABLE bloodtypes(
				id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
				name varchar(128) NOT NULL,
				description varchar(512) NOT NULL DEFAULT '',
				bonuses text[] DEFAULT NULL
			);`},
		[]string{ //v12
			`CREATE TABLE hunts(
				id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
				regionid INTEGER REFERENCES regions(id) NOT NULL,
				location varchar(128) NOT NULL,
				maxservants INTEGER NOT NULL DEFAULT 1,
				difficulty INTEGER NOT NULL DEFAULT 1,
				classbonus INTEGER REFERENCES bloodtypes(id)
			);`,
			`CREATE TABLE hunttimes(
				id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
				time INTEGER NOT NULL UNIQUE);`,
			`CREATE TABLE huntmodifiers(
				id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
				servantsamount INTEGER NOT NULL DEFAULT 1,
				bonus varchar(128) NOT NULL DEFAULT '');`,
			`CREATE TABLE huntrewardvalues(
				huntid INTEGER NOT NULL REFERENCES hunts(id),
				itemid INTEGER NOT NULL REFERENCES items(id),
				time INTEGER NOT NULL REFERENCES hunttimes(id), 
				modifier INTEGER NOT NULL REFERENCES huntmodifiers(id),
				minamount INTEGER NOT NULL,
				maxamount INTEGER NOT NULL,
				PRIMARY KEY(huntid, itemid, time, modifier));`},
		[]string{ //v13
			`CREATE TABLE factions(
				id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
			    name varchar(64) NOT NULL UNIQUE);`,
			`ALTER TABLE monsters
			ADD COLUMN bloodtypeid INTEGER NOT NULL REFERENCES bloodtypes(id);`,
			`ALTER TABLE monsters
			ADD COLUMN level INTEGER NOT NULL DEFAULT 1;`,
			`ALTER TABLE monsters
			ADD COLUMN factionid INTEGER NOT NULL REFERENCES factions(id);`,
			`ALTER TABLE monsters
			DROP COLUMN isboss;`,
			`ALTER TABLE monsters
			ADD COLUMN knowledgeid INTEGER REFERENCES knowledges(id) DEFAULT NULL;`,
			`ALTER TABLE monsters
			ADD COLUMN mapgenieid INTEGER DEFAULT NULL;`,
			`ALTER TABLE monsterdrops
			ADD COLUMN droprate REAL NOT NULL`,
			`ALTER TABLE monsterdrops
			ADD COLUMN amount INTEGER NOT NULL DEFAULT 1;`},
		[]string{ //v14
			`ALTER TABLE spells
			ALTER COLUMN description TYPE varchar(2048);`},
		[]string{ //v15
			`ALTER TABLE items
			ADD COLUMN maxstack INTEGER DEFAULT NULL;`},
		[]string{ //v16
			`ALTER TABLE spells
			ALTER COLUMN description TYPE text;`,
			`ALTER TABLE items
			ALTER COLUMN description TYPE text;`},
		[]string{ //v17
			`ALTER TABLE stations
			ADD COLUMN itemid INTEGER REFERENCES items(id) DEFAULT NULL;`}}
)
