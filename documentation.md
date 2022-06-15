- [knowledges](#knowledges)
- [itemtypes](#itemtypes)
- [sets](#sets)
- [items](#items)
- [tags](#tags)
- [itemtags](#itemtags)
- [itemstats](#itemstats)


### knowledges
Отвечает за связку источник знания - знание.
Знанием являются рецепты (recipe) и заклинания (spells). Источниками являются задания (quests), монстры с типом крови v blood и предметы с типом blueprint.
- id — связующий id, указывается у других сущностей.
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
