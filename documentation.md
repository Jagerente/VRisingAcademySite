- [knowledges](#knowledges)
- [itemtypes](#itemtypes)
- [sets](#sets)
- [items](#items)


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
- type — id itemtypes (оружие/броня/ювелирка/структура и т.п.)
- [knowledgid](#knowledges) — если этот предмет что-либо открывает; есть только у предметов типа blueprint
- setid — id сета, служит для группировки
- maxStack — макс. число предмета в одном стаке в инвентаре
```
insert into items(name, description, tier, type, knowledgeid, setid, maxStack) values
('Bone Sword', 'An all-round weapon that...', 3, 1, NULL, 1, 1)
```


