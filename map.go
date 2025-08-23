package main

type GameMap struct {
	Dungeons []Dungeon
}

func NewGameMap() GameMap {
	var dungeons []Dungeon
	var levels []Level
	l := NewLevel()
	levels = append(levels, l)

	d := Dungeon{
		Name:   "default",
		Levels: levels,
	}

	dungeons = append(dungeons, d)

	return GameMap{
		Dungeons: dungeons,
	}
}
