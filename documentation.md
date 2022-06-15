- [knowledges](#knowledges)
- [itemtypes](#itemtypes)
- [sets](#sets)
- [items](#items)
- [tags](#tags)
- [itemtags](#itemtags)
- [itemstats](#itemstats)
- [secondarystats](#secondarystats)
- [secondaryitemstats](#secondaryitemstats)
- [recipes](#recipes)
- [stations](#stations)
- [recipestations](#recipestations)
- [recipeingredients](#recipeingredients)
- [reciperesults](#reciperesults)


### knowledges
Отвечает за связку источник знания - знание.
Знанием являются рецепты ([recipe](#recipes)) и заклинания (spells). Источниками являются задания (quests), монстры с типом крови v blood и [предметы](#items) с типом blueprint.
- id — связующий id, указывается у других сущностей в поле knowledgeid.
```
insert into knowledges(id) values
(1)
```


### itemtypes
![image](https://user-images.githubusercontent.com/30572380/173904246-873fe3fb-4f85-4300-ad10-4c26ecca1d44.png)
- id
- title
```
insert into itemtypes(id, title) values
(1, 'Weapons')
```

### sets
![image](https://user-images.githubusercontent.com/30572380/173906506-f3c4d1c1-af6b-4995-a3bf-b04720c6ebeb.png)
- id
- name
- description — в основном NULL, используется только в сетах 2/4 предметов брони.
```
insert into sets(name, description) values
('Boneguard Vestment', '(2) Increase Max Health by 8.\n(4) Increase Physical Power by 2.')
```

### items
- id
- name — имя предмета
- description — описание
- tier — число, служит для сортировки
- [type](#itemtypes) — id itemtypes (оружие/броня/ювелирка/структура и т.п.)
- [knowledgid](#knowledges) — если этот предмет что-либо открывает; есть только у предметов типа blueprint
- [setid](#sets) — id сета, служит для группировки
- maxStack — макс. число предмета в одном стаке в инвентаре
```
insert into items(name, description, tier, type, knowledgeid, setid, maxStack) values
('Bone Sword', 'An all-round weapon that...', 3, 1, NULL, 1, 1)
```

### tags
Список всевозможных тегов.
- id
- value — имя тега
```
insert into tags (value) values
('equippable')
```

### itemtags
- [itemid](#items) — id предмета, у которого этот тег
- [tagid](#tags) — id тега

_Добавление предмету Bone Sword (id = 1) тегов equippable,blood-bound,weapon,sword,character menu_
```
insert into itemtags(itemid, tagid) values
(1, 2), (1, 3), (1, 4), (1, 5), (1, 80)
```

### itemstats
Описывает доп. хар-ки у экипируемых предметов (Weapon, Armour, Magic, Hat, Cloak)
- id
- mainstat — число;
  - для предметов типа Weapon — это Physical Power,
  - для предметов типа Armour и Cloak — это Max Health,
  - для предметов типа Magic — это Spell Power.
- setbonus — строка; НЕ ИСПОЛЬЗУЕТСЯ
- gearlevel — есть только у Weapon, Armour и Magic.
- durability — есть только у Weapon, Armour и Magic.
- slotid — на данный момент отдельной таблицы НЕТ, захардкодено.
  - hat         1
  - chestguard	2
  - leggings	  3
  - gloves	    4
  - boots	      5
  - cloak	      6
  - jewellery	  7
  - no	        0

_Пример для Bone Sword с физ. уроном 5.2, gear level 3, прочностью 817_
```
insert into itemstats(id, mainstat, setbonus, gearlevel, durability, slotid) values
(1, 5.2, '', 3, 817, 0)
```

### secondarystats
Список всевозможных доп. бонусов у предметов (Weapon, Armour, Magic, Cloak)
- id
- bonus
```
insert into secondarystats(bonus) values
('+25% Physical Damage to Vegetation')
```

### secondaryitemstats
- [statsid](#itemstats) — id itemstats
- [secondarystatid](#secondarystats)

_Пример добавления Bone Sword, к которому привязан itemstats id = 1, способности +25% Physical Damage to Vegetation_
```
insert into secondaryitemstats(statsid, secondarystatid) values
(1, 1)
```

### recipes
![image](https://user-images.githubusercontent.com/30572380/173920337-8c742320-d5b5-4a0c-a785-f8707b2e1c7f.png)

- id
- time — время крафта (без учёта скидки 20% за закрытую комнату)
- [knowledgeid](#knowledges) — чем открывается; NULL, если открыт по умолчанию.
_Vermin Salve крафтится 10 секунд, и открывается после выполнения задания, чей knowledgeid = 3_
```
insert into recipes(time, knowledgeid) values
(10, 3)
```

### stations
Станции, на которых могут изгатавливаться рецепты.
- id
- name
- description

Список станций, которые могут предоставить скидки за Matching Floor (-25% к ресурсам) или за Confined Room (-20% к времени):
_simple workbench,sawmill,furnace,grinder,tannery,blood press,woodworking bench,vermin nest,alchemy table,tailoring bench,smithy,loom,jewelcrafting table,gem cutting table,paper press,anvil_
```
insert into stations(name, description) values
('Character Menu', 'Open your character menu to craft some basic survival tools.')
```

### recipestations
![image](https://user-images.githubusercontent.com/30572380/173920487-01e24456-bf5a-4721-8fb5-d7e10db34644.png)
- [recipeId](#recipes)
- [stationId](#stations)

_Пример Vermin Salve (recipe id = 10), который крафтится на станциях Character Menu (id = 1) и Alchemy Table (id = 11)_
```
insert into recipestations(recipeId, stationId) values
(10, 1), (10, 11)
```

### recipeingredients
![image](https://user-images.githubusercontent.com/30572380/173920545-4ddcf8d4-c15e-473a-934a-ddc45233d640.png)
- [recipeId](#recipes)
- [itemId](#items)
- amount

_Vermin Salve (recipe id = 10) крафтится с использованием 1 шт. Rat (id = 157), 60 шт. Plant Fibre (id = 187) и 20 шт. Bone (id = 20)_
```
insert into recipeingredients(recipeId, itemId, amount) values
(10, 157, 1), (10, 187, 60), (10, 182, 20)
```

### reciperesults
![image](https://user-images.githubusercontent.com/30572380/173920577-561c5de3-6002-430f-925d-754d4df3eb23.png)
- [recipeId](#recipes)
- [itemId](#items)
- amount

_В результате крафта Vermin Salve (recipe id = 10) получается... Vermin Salve (id = 163) в кол-ве 1 шт._
```
insert into reciperesults(recipeId, itemId, amount) values
(10, 163, 1)
```
